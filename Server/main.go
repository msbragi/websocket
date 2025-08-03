package main

import (
	"log"
	"net/http"
	"os"
)

const (
	staticDir  = "./static"
	serverPort = "8080"
)

func main() {
	// Ensure static directory exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Fatalf("Static directory %s does not exist", staticDir)
	}

	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	log.Printf("Static server running on http://localhost:%s/", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
