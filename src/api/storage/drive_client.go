package storage

import (
	"context"
	"errors"
	"os"

	"cefetdb2api/config"
	"cefetdb2api/storage/auth"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type DriveClient struct{}

func NewDriveClient() *DriveClient {
	return &DriveClient{}
}

func (c *DriveClient) GetDriveService(ctx context.Context) (*drive.Service, error) {
	cfg := config.NewConfig()
	cfg.ReadConfigFile()
	credentialsFile := cfg.DriveAuth.CredentialsFilePath
	tokenFile := cfg.DriveAuth.TokenPath

	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, errors.New("unable to read drive client credendials file")
	}

	// If modifying these scopes, delete your previously saved token.json.
	driveConfig, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		return nil, err
	}

	client, err := auth.GetClient(ctx, driveConfig, tokenFile)
	if err != nil {
		return nil, err
	}

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return srv, err
}
