package types

import (
	"encoding/json"
	"errors"
	"os"
)

type OAuthCredentials struct {
	Installed struct {
		ClientID                string   `json:"client_id"`
		ProjectID               string   `json:"project_id"`
		AuthURI                 string   `json:"auth_uri"`
		TokenURI                string   `json:"token_uri"`
		AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
		ClientSecret            string   `json:"client_secret"`
		RedirectUris            []string `json:"redirect_uris"`
	} `json:"installed"`
}

func NewOAuthCredentials() OAuthCredentials {
	return OAuthCredentials{}
}

func (c *OAuthCredentials) GetCredentialsFromFile(filePath string) (OAuthCredentials, error) {

	file, err := os.ReadFile(filePath)
	if err != nil {
		return OAuthCredentials{}, errors.New("unable to read OAuth credentials file")
	}

	json.Unmarshal(file, &c)

	return *c, nil
}
