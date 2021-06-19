package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gitlab.com/rayquen-google/golang/auth/auth_service"
	"gitlab.com/rayquen-google/golang/auth/auth_service_spreadsheet"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {

	var auth auth_service.IAuthService = &auth_service_spreadsheet.AuthServiceSpreadsheet{}

	err := auth.Initialize("P:/rayquen-google/golang/auth/credentials.json", "P:/rayquen-google/golang/auth/token.json", true)
	if err != nil {
		log.Fatalf("[Main] Error al inicializar %v", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "--request-token" {

		fmt.Println("Solicitando Token...")
		err = auth.RequestToken()
		if err != nil {
			log.Fatalf("[Main] Error al solicitar token %v", err)
		}

	} else {
		fmt.Println("Autenticando...")
		err = auth.Authenticate()
		if err != nil {
			log.Fatalf("[Main] Error al autenticar %v", err)
		}
	}

	fmt.Println("Ejecutando...")

	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(auth.GetClient()))
	if err != nil {
		log.Fatalf("[Main::NewService] %v", err)
	}

	var spreadsheet_id = "1BPGEDtDsiHKNfJylUFfEy9esnYY1If6SAKHW82psthA"
	var spreadsheet_page = "Clientes"

	readRange := spreadsheet_page + "!A1:I24"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheet_id, readRange).Do()

	if err != nil {
		log.Fatalf("[Main::GetValues] %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("Nada!!")
	} else {
		for _, row := range resp.Values {
			fmt.Printf("%s: %s\n", row[0], row[2])
		}
	}
}
