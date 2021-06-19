package auth_service_spreadsheet

import (
	"log"

	"gitlab.com/rayquen-google/golang/auth/auth_service"
)

type AuthServiceSpreadsheet struct {
	auth_service.AuthService
}

func (s *AuthServiceSpreadsheet) Initialize(credentialFile string, tokenFile string, readonly bool) error {
	err := s.Credential.Load(credentialFile)
	if err != nil {
		log.Printf("[AuthServiceSpreadsheets::Initialize] con %v %v", credentialFile, err)
		return err
	}

	err = s.Token.Load(tokenFile)
	if err != nil {
		log.Printf("[AuthServiceSpreadsheets::Initialize] con %v %v", tokenFile, err)
		return err
	}

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
