package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

const ConfigFile = "./config/config.json"

type Config struct {
	ServerPort string `json:"server_port"`
	UploadDir  string `json:"upload_dir"`
	UpdateUrl  string `json:"update_url"`
	WsStatus   string `json:"ws_status"`
}

func handleConfig(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Config read error:", err)
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			sendError(conn, "Invalid config message")
			continue
		}

		switch msg.Type {
		case "getconfig":
			getConfig(conn, ConfigFile)
		case "setconfig":
			saveConfig(conn, ConfigFile, msg.Config)
		default:
			sendError(conn, "Unknown config message type")
		}
	}
}

func getConfig(conn *websocket.Conn, path string) {
	cfg, err := loadConfig(path)
	if err != nil {
		sendError(conn, err.Error())
		return
	}
	jsonCfg, err := json.Marshal(cfg)
	if err != nil {
		sendError(conn, "Failed to encode config")
		return
	}
	response := Message{
		Type:   "configdata",
		Config: jsonCfg,
	}
	sendJSON(conn, response)
}

func saveConfig(conn *websocket.Conn, path string, raw json.RawMessage) {
	var newCfg Config
	if err := json.Unmarshal(raw, &newCfg); err != nil {
		sendError(conn, "Invalid config data")
		return
	}
	f, err := os.Create(path)
	if err != nil {
		sendError(conn, "Failed to create config file")
		return
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(newCfg); err != nil {
		sendError(conn, "Failed to write config file")
		return
	}
	sendSuccess(conn, "Config updated successfully")
}

func loadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
