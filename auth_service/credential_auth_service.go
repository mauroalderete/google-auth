package auth_service

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Estructura basica para operar con el archivo de credenciales
type CredentialAuthService struct {
	path   string         // ruta y nombre del archivo de credencial
	data   []byte         // contenido del archivo de credencial
	err    error          // descripcion del error en caso de que no se haya podido cargar la credencial
	config *oauth2.Config // resultado de la configuracion una vez solicitada usando la credencial
}

// Inicializa la ruta del archivo de credencial y gestiona la carga de la credencial
// De esta forma se consige una validación del contenido del archivo
func (c *CredentialAuthService) Load(path string) error {
	c.path = filepath.Clean(path)

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

// Parsea el archivo de credencial a la estructura *oauth2.Config
// utilizada por google para gestionar una autenticacion o un token nuevo
func (c *CredentialAuthService) GetConfig(fromUrlService string) error {

	config, err := google.ConfigFromJSON(c.data, fromUrlService)
	if err != nil {
		log.Printf("[CredentialAuthService::GetConfig] %v", err)
		return err
	}

	c.config = config

	return nil
}
