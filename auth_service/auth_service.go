package auth_service

import (
	"context"
	"log"
	"net/http"
)

// IAuthService expone los metodos basicos para implementar la autenticacion
// Se recomienda usar la interfaz para evitar acceder a funciones y metodos
// que pueden quedar expuestos en las estructuras que la implementan
// Posee los metodos escenciales de inicializacion, autenticacion y solicitud de token
// Un metodo extra permite acceder al *http.Client para solicitar los servicios
type IAuthService interface {
	Initialize(credentialFile string, tokenFile string, readonly bool) error

	Authenticate() error
	RequestToken() error
	GetClient() *http.Client
}

// AuthService es la estructura base de un servicio de autenticacion
// Incorpora un gestor de token, credencial y el modo de acceso
// mantiene una copia de la url utilizada por la api y un *http.Client
type AuthService struct {
	Token      TokenAuthService      // gestiona el token generado tras una solicitud de token
	Credential CredentialAuthService // gestiona el acceso y operaciones basados en las credenciales
	Readonly   bool                  // indica si se consumira un servicio en modo letura o lectura/escritura
	Url        string                // mantiene la url del servicio a consumir, depende de Readonly
	Client     *http.Client          // mantiene un cliente http sobre el cual se solicitaran los servicios a consumir
}

// Initialize permite realizar una inicializacion basica y comun a todos los tipos de autenticacion
// recibe dos rutas de archivos. Si bien no es necesario que tokenFile exista
// para ejecutar otras operaciones como RequestToken, es recomendable
// utilizar una ubicación real, ya que en caso de persistir el token
// se realizara en esa ruta
func (s *AuthService) Initialize(credentialFile string, tokenFile string) (error, error) {

	errCredential := s.Credential.Load(credentialFile)
	if errCredential != nil {
		log.Printf("[AuthServiceSpreadsheets::Initialize] con %v %v", credentialFile, errCredential)
	}

	errToken := s.Token.Load(tokenFile)
	if errToken != nil {
		log.Printf("[AuthServiceSpreadsheets::Initialize] con %v %v", tokenFile, errToken)
	}

	return errCredential, errToken
}

// Authenticate gestiona el procedimiento de autenticación
// Se debe considerar que tanto la credencial como el token
// hayan sido inicializados correctamente antes de comenzar con la operación de autenticación.
// El proceso termina obteniendo una instancia de un cliente http
// Esta instancia no es validada, por lo que se delega la responsabilidad
// de verificar que el proceso de autenticación haya resultado exitosamente
func (s *AuthService) Authenticate() error {

	//valido que tenga una credencial inicializada
	if s.Credential.err != nil {
		log.Println("[AuthService::Authenticate] La credencial no estan lista")
		return s.Credential.err
	}

	//valido que tenga un token inicializado
	if s.Token.err != nil {
		log.Println("[AuthService::Authenticate] El token no esta listo")
		return s.Token.err
	}

	//obtengo un cliente a partir de la configuracion de credencial inicializada

	s.Client = s.Credential.config.Client(context.Background(), s.Token.token)

	return nil
}

// RequestToken solicita un token nuevo al servicio de google.
// Requiere que exista una credencial valida
// Debido a que el metodo de verificación de permiso necesita que un usuario intractute con el prompt
// no es recomendable para usar en sistema automatizados
// Al finalizar la operación persiste el token en la ruta indicada durante la inicialización
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

// GetClient devuelve el puntero al cliente http.
// No se valida que el cliente este instanciado o sea funcional
func (s *AuthService) GetClient() *http.Client {
	return s.Client
}
