package auth

import (
	"context"
	"encoding/json"
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
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
