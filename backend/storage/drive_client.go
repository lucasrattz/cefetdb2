package storage

import (
	"context"
	"encoding/json"
	"errors"

	"cefetdb2api/config"
	"cefetdb2api/storage/auth"
	"cefetdb2api/types"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type DriveClient struct {
	credentials types.OAuthCredentials
}

func NewDriveClient(c types.OAuthCredentials) *DriveClient {
	return &DriveClient{c}
}

func (c *DriveClient) GetDriveService(ctx context.Context, cfg *config.Config) (*drive.Service, error) {
	tokenFile := cfg.DriveAuth.TokenPath

	b, err := json.Marshal(c.credentials)
	if err != nil {
		return nil, errors.New("unable to marshal drive client credendials")
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
