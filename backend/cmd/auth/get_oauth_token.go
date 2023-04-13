package main

import (
	"cefetdb2api/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
)

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(credentialsPath string) (*oauth2.Token, error) {
	ctx := context.Background()

	b, err := os.ReadFile(credentialsPath)
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, also modify in api/storage/drive_client.go
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		return nil, err
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, err
	}

	tok, err := config.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}

	return tok, nil
}

// Saves a token to a file path.
func saveToken(tokenPath string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", tokenPath)
	f, err := os.OpenFile(tokenPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}

func main() {
	cfg := config.NewConfig()
	cfg, err := cfg.ReadConfigFile()
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	tok, err := getTokenFromWeb(cfg.DriveAuth.CredentialsFilePath)
	if err != nil {
		log.Fatalf("Unable to get token from web: %v", err)
	}

	err = saveToken(cfg.DriveAuth.TokenPath, tok)
	if err != nil {
		log.Fatalf("Unable to save token: %v", err)
	}
}
