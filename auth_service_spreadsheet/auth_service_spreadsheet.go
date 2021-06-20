package auth_service_spreadsheet

import (
	"log"

	"gitlab.com/rayquen-google/golang/auth/auth_service"
)

// AuthServiceSpreadsheet es una especialización de AuthService
// Se compone de los mismos campos y estructuras
// Permite construir comportamientos especificos
type AuthServiceSpreadsheet struct {
	auth_service.AuthService
}

// Amplia la inicializacion de AuthService agregando el modo de acceso
// y especificando la Url del servicio a consumir
// Efectua la solicitud para obtener la configuración de la autenticación
func (s *AuthServiceSpreadsheet) Initialize(credentialFile string, tokenFile string, readonly bool) error {

	var base auth_service.AuthService
	err := base.Initialize(credentialFile, tokenFile)
	if err != nil {
		log.Printf("[AuthServiceSpreadsheets::Initialize] Base %v", err)
		return err
	}

	s.Credential = base.Credential
	s.Token = base.Token

	s.Readonly = readonly
	if s.Readonly {
		s.Url = "https://www.googleapis.com/auth/spreadsheets.readonly"
	} else {
		s.Url = "https://www.googleapis.com/auth/spreadsheets"
	}

	err = s.Credential.GetConfig(s.Url)
	if err != nil {
		log.Printf("[AuthServiceSpreadsheets::Initialize] %v", err)
		return err
	}

	return nil
}
