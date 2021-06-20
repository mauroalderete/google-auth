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

// Estructura basica para operar con el archivo de token
type TokenAuthService struct {
	path  string        // ruta y nombre del archivo que contiene el token
	err   error         // descripcion del error en caso de que no se haya podido cargar el token
	token *oauth2.Token // resultado del parseo del token
}

// Carga el archivo de token y decodifica el contenido
// El resultado se almacena en token como entidad *oauth2.token
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

// Permite obtener un token nuevo utilizando los mecanismos de verificación y permisos de google
// Requiere acceso al prompt para poder registrar un hash de confirmación
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

// Persiste en la ruta path el token cargado
// Puede ser el mismo token leido al inicializar
// o el token recibido al solicitar uno nuevo
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
