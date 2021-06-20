package auth_service

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

var isCredentialLoadableTest = []struct {
	path     string
	content  string
	expected bool
}{
	{"./credential.json.test", "abcdef", true},
	{"./otro", "abcdef", false},
}

func TestRead(t *testing.T) {
	var c CredentialAuthService

	err := ioutil.WriteFile(
		isCredentialLoadableTest[0].path,
		[]byte(isCredentialLoadableTest[0].content),
		fs.ModeAppend.Perm())

	if err != nil {
		t.Errorf("No se pudo preparar el archivo para las pruebas")
	}

	for _, tt := range isCredentialLoadableTest {
		t.Run(tt.path, func(t *testing.T) {
			got := c.Load(tt.path)
			if (got == nil) != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}

	err = os.Remove(isCredentialLoadableTest[0].path)
	if err != nil {
		fmt.Printf("no se pudo limpiar la prueba al finalizar %v", err)
	}
}
