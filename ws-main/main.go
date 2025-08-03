package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

func main() {
	// Load config
	cfg, err := loadConfig(ConfigFile)
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(cfg.UploadDir, 0755); err != nil {
		log.Fatal("Failed to create uploads directory:", err)
	}

	/*
	* Handle file upload download
	 */
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		handleFiles(w, r, cfg.UploadDir)
	})

	/*
	* Handle config
	 */
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		handleConfig(w, r)
	})

	/*
	* Handle browser
	 */
	http.HandleFunc("/browse", func(w http.ResponseWriter, r *http.Request) {
		handleBrowser(w, r, cfg.UploadDir)
	})

	log.Println("Server starting on :" + cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
