package types

import "time"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Expiry       string `json:"expiry"`
}

func (t Token) GetExpiryInTime() (time.Time, error) {
	expiry, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", t.Expiry)
	if err != nil {
		return time.Time{}, err
	}

	return expiry, nil
}
