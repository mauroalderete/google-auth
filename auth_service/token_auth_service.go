package auth_service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

type TokenAuthService struct {
	path  string
	err   error
	token *oauth2.Token
}

func (t *TokenAuthService) Load(path string) error {
	t.path = filepath.Clean(path)

	f, err := os.Open(t.path)
	if err != nil {
		log.Printf("[TokenAuthService::loadToken] open %v", err)
		return err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		log.Printf("[TokenAuthService::loadToken] decode %v", err)
		return err
	}
	t.token = tok
	return nil
}

func (t *TokenAuthService) GetTokenFromWeb(config *oauth2.Config) error {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return fmt.Errorf("unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web: %v", err)
	}

	t.token = tok

	return nil
}

func (t *TokenAuthService) SaveTokenFile() error {
	fmt.Printf("Saving credential file to: %s\n", t.path)
	f, err := os.OpenFile(t.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(t.token)

	return nil
}
