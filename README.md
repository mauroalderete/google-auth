# auth

Libreria que gestiona la autenticación a la apiv4 de google

## Resumen

Permite obtener un *http.Client listo para usar y solicitar un servicio de google por medio de su api.

Tambien permite gestionar un token nuevo.

En todos los casos se requiere una credencial valida para iniciar la conexion. La credencial se puede obtener desde la consola de google.

## Uso

Se debe instanciar el servicio que se desea utilizar por medio de la interfaz IAuthService. Por ejemplo, para obtener el permiso para usar las hojas de calculo de google se podria implementar de la siguiente forma

```go
var auth auth_service.IAuthService = &auth_service_spreadsheet.AuthServiceSpreadsheet{}
```

Una vez instanciado el servicio de autorización, es necesario inicializar el servicio con las rutas de las credenciales, tokens y el modo de consumo (lectura o lectura/escritura)

```go
err := auth.Initialize("./credentials.json", "./token.json", true)
if err != nil {
	log.Fatalf("[Main] Error al inicializar %v", err)
}
```

En el ejemplo se intenta inicializar una autenticación del servicio de hojas de calculo en modo lectura utilizando los archivos correspondiente. Si alguno de los archivos no puede ser encontrado, o no tiene el formato correcto se emitira un error. Correspondera a otra instancia evaluar como operar en estos casos

Una vez inicializado, se puede optar por ejecutar la autenticación o generar un token nuevo.

La autenticación consiste en ejecutar la función correspondiente

```go
err = auth.Authenticate()
if err != nil {
	log.Fatalf("[Main] Error al autenticar %v", err)
}
```

Se devolvera error si la credencial o token no estan bien inicializadas, o al tratar de obtener la configuración *oauth.Config

Por el momento no se realiza una validación del resultado de la autenticación. Por lo que se debera efectuar esta operación mas adelante.

Un caso típico de este problema se da cuando el token esta vencido. La autenticación finaliza su proceso correctamente, pero al momento de solicitar un servicio o tratar de operar con él, se emiten mensajes de errores indicando que el token esta vencido.

En este caso se puede ejecutar la solicitud de un token nuevo.

```go
err = auth.RequestToken()
if err != nil {
	log.Fatalf("[Main] Error al solicitar token %v", err)
}

err := auth.Initialize("./credentials.json", "./token.json", true)
if err != nil {
	log.Fatalf("[Main] Error al inicializar %v", err)
}

err = auth.Authenticate()
if err != nil {
	log.Fatalf("[Main] Error al autenticar %v", err)
}
```

La solicitud de un token hace uso de la interfaz de consola, por lo que no es recomendable utilizar este método en un ecosistema de servidor sin acceso a la interfaz, o donde el usuario no pueda interactuar con el prompt.
Al ejecutar la función se le mostrará al usuario una url que deberá abrir en un navegador. Al aceptar los permisos se le otorgará un hash que tendrá que insertar en el prompt. Si todo sale bien, se grabara el token recibido en la misma ruta que se utilizo al indicar la inicializacion.

Luego para poder utilizar el nuevo token es necesario volver a inicializar y autenticar la conexión como se hace en el ejemplo.

