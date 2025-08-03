package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/websocket"
)

func handleFiles(w http.ResponseWriter, r *http.Request, uploadDir string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Files handler read error:", err)
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("JSON error:", err)
			continue
		}
		switch msg.Type {
		case "getfile":
			handleGetFile(conn, msg.Filename, uploadDir)
		case "putfile":
			handlePutFile(conn, msg.Filename, msg.Data, uploadDir)
		default:
			sendError(conn, "Unknown message type")
		}
	}
}

func handleGetFile(conn *websocket.Conn, filename string, uploadDir string) {
	if filename == "" {
		sendError(conn, "Filename required")
		return
	}

	filePath := filepath.Join(uploadDir, filename)
	data, err := os.ReadFile(filePath)
	if err != nil {
		sendError(conn, fmt.Sprintf("File read error: %v", err))
		return
	}

	response := Message{
		Type:     "filedata",
		Filename: filename,
		Data:     data,
	}

	if err := sendJSON(conn, response); err != nil {
		log.Println("Send error:", err)
	}
}

func handlePutFile(conn *websocket.Conn, filename string, data []byte, uploadDir string) {
	if filename == "" {
		sendError(conn, "Filename required")
		return
	}

	filePath := filepath.Join(uploadDir, filename)
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		sendError(conn, fmt.Sprintf("File write error: %v", err))
		return
	}

	response := Message{
		Type:     "success",
		Filename: filename,
		Message:  "File saved successfully",
	}

	if err := sendJSON(conn, response); err != nil {
		log.Println("Send error:", err)
	}
}
