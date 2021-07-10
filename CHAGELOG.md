# CHANGELOG.md

## 10/07/2021 Hotfix v1.0.1

Cuando se solicitaba el token a partir de las credenciales existentes, si el archivo de destino, token.json por ejemplo, no existe entonces la referencia al objeto config de las credenciales se mantenia nula. Esto se daba a que el servicio de autorizacion de spreadsheet nunca llegaba a parsear la configuración de las credenciales.

Esto se solucionó modificando el método base de inicialización para que retorne dos clases de errores, cada uno correspondiente a cada archivo.

De esta forma, si la inicialización produjo un error únicamente al procesar el archivo de token inaccesible, pero no con las credenciales, el parseo de las credenciales se lleva a cabo de todas maneras. Al final la inicialización de AuthServiceSpreadsheet seguira dando error por el token.

Con esto se permite la instanciación de un objeto *oauth2.Config a partir de las credenciales. Lo que permite ser consultado al solicitar un toquen nuevo si es necesario.

## 20/06/2021 Release v1.0.0

Libero primera version que incluye los servicios de authenticacion a apis de google
Por el momento solo es posible conectar a la api de spreadsheets

Se documenta caso de uso

## 20/06/2021 Feature refactoring

Descarto codigo en desuso
Renombro metodos
Genero pruebas unitarias simples, sin emulacion de entorno
Retiro codigo de ejemplo

## 19/06/2021 Feature Import codebase

Resuelvo el problema de packetes e importaciones que tenia el proyecto anterior: https://gitlab.com/vyra/vyra-database-migration

Construyo entidades para gestionar los diferentes aspectos de la autorizacion:

- Credenciales:
    - gestiona la lectura del archivo
    - carga en memoria de los datos de la credencial
    - cofiguracion de oauth2
- Token:
    - Carga en memoria el token persistido
    - Solicita un nuevo token al usuario
    - Persiste el nuevo token recibido
- AuthService:
    - Es una entidad base de la cual heredan las entidades de autorizacion especializadas
    - Gestiona la inicializacion y verificacion de credenciales y token.
    - Gestiona la solicitud de un nuevo token y su persistencia
    - Implementa una interfaz para exportar las funcionalidades
- Main:
    - Autentica o solicita un nuevo token dependiendo del flag ```--request-token```
    - No termina de verificar si tiene acceso o no
    - Trata de conectarse iniciando una nueva instancia de un servicio y ejecuta una consulta

## 19/06/2021 Feature Init

Inicializo los archivos, ignores, documentos y base del proyecto