package main

import (
	"log"

	"cefetdb2api/config"
	api "cefetdb2api/server"
	"cefetdb2api/types"
)

func main() {
	cfg := config.NewConfig()
	cfg, err := cfg.ReadConfigFile()
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	port := cfg.Server.Port
	credentialsFilePath := cfg.DriveAuth.CredentialsFilePath

	c := types.NewOAuthCredentials()
	c, err = c.GetCredentialsFromFile(credentialsFilePath)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(port)
	log.Printf("Starting server on port %s", port)
	log.Fatal(server.Start())
}
