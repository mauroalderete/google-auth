package auth_service

import (
	"context"
	"log"
	"net/http"
)

type IAuthService interface {
	Initialize(credentialFile string, tokenFile string, readonly bool) error

	Authenticate() error
	RequestToken() error
	GetClient() *http.Client
}

type AuthService struct {
	Token      TokenAuthService
	Credential CredentialAuthService
	Readonly   bool
	Url        string
	Client     *http.Client
}

func (s *AuthService) Authenticate() error {

	//valido que tenga una credencial inicializada
	if s.Credential.err != nil {
		log.Println("[AuthService::Authenticate] La credencial no estan lista")
		return s.Credential.err
	}

	//valido que tenga un token inicializado
	if s.Token.Error() != nil {
		log.Println("[AuthService::Authenticate] El token no esta listo")
		return s.Token.Error()
	}

	//obtengo un cliente a partir de la configuracion de credencial inicializada

	s.Client = s.Credential.config.Client(context.Background(), s.Token.token)

	return nil
}

func (s *AuthService) RequestToken() error {

	//valido que tenga una credencial inicializada
	if s.Credential.err != nil {
		log.Println("[AuthService::Authenticate] La credencial no estan lista")
		return s.Credential.err
	}

	err := s.Token.GetTokenFromWeb(s.Credential.config)
	if err != nil {
		return err
	}

	err = s.Token.SaveTokenFile()
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GetClient() *http.Client {
	return s.Client
}
