package util

import (
	"context"
	"errors"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetAccessToken(refreshToken string) (string, error) {

	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}

	ctx := context.Background()
	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	newToken, err := config.TokenSource(ctx, token).Token()
	if err != nil {
		log.Print(err)
		return "", errors.New("[ERROR] Generate Access Token Error!")
	}

	return newToken.AccessToken, nil
}
