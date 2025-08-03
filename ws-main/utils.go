package main

import (
	"bytes"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type     string          `json:"type"`
	Filename string          `json:"filename,omitempty"`
	Data     []byte          `json:"data,omitempty"`
	Config   json.RawMessage `json:"config,omitempty"`
	Message  string          `json:"message,omitempty"`
}

func sendMessage(conn *websocket.Conn, msgType string, msg string) {
	response := Message{
		Type:    msgType,
		Message: msg,
	}
	sendJSON(conn, response)
}

func sendError(conn *websocket.Conn, message string) {
	sendMessage(conn, "error", message)
}

func sendSuccess(conn *websocket.Conn, message string) {
	sendMessage(conn, "success", message)
}

func sendJSON(conn *websocket.Conn, msg any) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(msg); err != nil {
		return err
	}
	return conn.WriteMessage(websocket.TextMessage, buf.Bytes())
}
