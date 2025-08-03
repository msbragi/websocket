package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/websocket"
)

func handleBrowser(w http.ResponseWriter, r *http.Request, rootDir string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Browser read error:", err)
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			sendError(conn, "Invalid browser message")
			continue
		}

		switch msg.Type {
		case "listdir":
			handleListDir(conn, msg.Filename, rootDir)
		case "mkdir":
			handleMkdir(conn, msg.Filename, msg.Message, rootDir)
		case "rmdir":
			handleRmdir(conn, msg.Filename, msg.Message, rootDir)
		default:
			sendError(conn, "Unknown browser message type")
		}
	}
}

func handleListDir(conn *websocket.Conn, dir string, rootDir string) {
	if dir == "" {
		dir = rootDir
	}

	absRoot, _ := filepath.Abs(rootDir)
	absDir, _ := filepath.Abs(dir)
	if !strings.HasPrefix(absDir, absRoot) {
		absDir = absRoot
	}

	parent := filepath.Dir(absDir)
	if absRoot == absDir {
		parent = ""
	}

	entries, err := os.ReadDir(absDir)

	if err != nil {
		sendError(conn, "Failed to read directory")
		return
	}

	var folders []string
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}

	payload, _ := json.Marshal(map[string]any{
		"folders": folders,
		"path":    absDir,
		"parent":  parent,
	})

	response := Message{
		Type:   "dirlist",
		Config: payload,
	}

	sendJSON(conn, response)
}
func handleMkdir(conn *websocket.Conn, dir, folderName, rootDir string) {
	if dir == "" {
		dir = rootDir
	}
	newPath := filepath.Join(dir, folderName)
	if err := os.Mkdir(newPath, 0755); err != nil {
		sendError(conn, "Failed to create folder")
		return
	}
	handleListDir(conn, dir, rootDir)
}

func handleRmdir(conn *websocket.Conn, dir, folderName, rootDir string) {
	if dir == "" {
		dir = rootDir
	}
	targetPath := filepath.Join(dir, folderName)
	if err := os.Remove(targetPath); err != nil {
		sendError(conn, "Failed to delete folder")
		return
	}
	handleListDir(conn, dir, rootDir)
}
