package auth_service

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type CredentialAuthService struct {
	path   string
	data   []byte
	err    error
	config *oauth2.Config
}

func (c *CredentialAuthService) Load(path string) error {
	c.path = filepath.Clean(path)

	return c.Read()
}

func (c *CredentialAuthService) Read() error {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		log.Printf("[CredentialAuthService::Load] %v", err)
		c.data = nil
		c.err = err
		return err
	}
	c.data = make([]byte, len(data))
	copy(c.data, data)
	c.err = nil
	return nil
}

func (c *CredentialAuthService) Path() *string {
	return &c.path
}

func (c CredentialAuthService) Data() *[]byte {
	return &c.data
}

func (c *CredentialAuthService) Error() error {
	return c.err
}

func (c *CredentialAuthService) GetConfig(fromUrlService string) error {

	config, err := google.ConfigFromJSON(c.data, fromUrlService)
	if err != nil {
		log.Printf("[CredentialAuthService::GetConfig] %v con %v", err, c.data)
		return err
	}

	c.config = config

	return nil
}

func (c *CredentialAuthService) Config() *oauth2.Config {
	return c.config
}