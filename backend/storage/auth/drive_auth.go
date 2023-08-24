package auth

import (
	"cefetdb2api/types"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

// Retrieve a token, saves the token, then returns the generated client.
func GetClient(ctx context.Context, oauthConfig *oauth2.Config, tokenFile string) (*http.Client, error) {
	// Generate the token file with the oauth_token.go script under cmd/auth

	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		return nil, err
	}

	return oauthConfig.Client(ctx, tok), nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	byteValue, _ := io.ReadAll(f)
	var payload types.Token
	err = json.Unmarshal(byteValue, &payload)
	if err != nil {
		return nil, err
	}

	expiry, err := payload.GetExpiryInTime()
	if err != nil {
		return nil, err
	}

	tok := &oauth2.Token{}
	tok.AccessToken = payload.AccessToken
	tok.RefreshToken = payload.RefreshToken
	tok.TokenType = payload.TokenType
	tok.Expiry = expiry
	return tok, err
}
