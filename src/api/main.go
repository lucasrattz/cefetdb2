package main

import (
	"log"
	"os"

	"cefetdb2api/server"
)

func main() {
	port := "8080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting server on port %s", port)

	server := server.NewServer(port)
	log.Fatal(server.Start())
}
