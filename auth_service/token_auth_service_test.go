package auth_service

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

var isTokenLoadableTest = []struct {
	path     string
	content  string
	expected bool
}{
	{"./token.json.test", "{\"access_token\":\"abc\",\"token_type\":\"Bearer\",\"refresh_token\":\"abc\",\"expiry\":\"2021-06-19T21:15:54.8451088-03:00\"}", true},
	{"./otro", "", false},
}

func TestTokenAuthServiceLoad(t *testing.T) {
	var tok TokenAuthService

	err := ioutil.WriteFile(
		isTokenLoadableTest[0].path,
		[]byte(isTokenLoadableTest[0].content),
		fs.ModeAppend.Perm())

	if err != nil {
		t.Errorf("No se pudo preparar el archivo para las pruebas")
	}

	for _, tt := range isTokenLoadableTest {
		t.Run(tt.path, func(t *testing.T) {
			got := tok.Load(tt.path)
			if (got == nil) != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}

	err = os.Remove(isTokenLoadableTest[0].path)
	if err != nil {
		fmt.Printf("no se pudo limpiar la prueba al finalizar %v", err)
	}
}
