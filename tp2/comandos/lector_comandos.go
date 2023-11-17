package comandos

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAHash "tdas/diccionario"
	"tp2/errores"
	"tp2/usuarios"
)

// Comandos posibles:
const (
	COMANDO_LOGIN              = "login"
	COMANDO_LOGOUT             = "logout"
	COMANDO_PUBLICAR           = "publicar"
	COMANDO_VER_SIGUIENTE_FEED = "ver_siguiente_feed"
	COMANDO_LIKEAR_POST        = "likear_post"
	COMANDO_MOSTRAR_LIKES      = "mostrar_likes"
)

// ejemplos comandos
// login chicho1994
// logout
// publicar Tiene todo el dinero del mundo...
// ver_siguiente_feed
// likear_post 0
// mostrar_likes 0
type Comando int

const (
	LOGIN Comando = iota
	LOGOUT
	PUBLICAR
	VER_SIGUIENTE_FEED
	LIKEAR_POST
	MOSTRAR_LIKES
	ERROR
)

func LectorComandos(dicUsuario TDAHash.Diccionario[string, usuarios.Usuario]) {
	var err error

	var conectado bool
	var nombre string
	var ptrnombre *string
	ptrnombre = &nombre
	var usuarioLogueado usuarios.Usuario
	var ptrUsuarioLogueado *usuarios.Usuario
	ptrUsuarioLogueado = &usuarioLogueado
	conectado = false
	linea := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		// posible error
		linea, err = reader.ReadString('\n')
		if err != nil {
			// Verifica si se llego al final del archivo
			if err.Error() == "EOF" {
				break
			}

		}

		entradaSeparada := separarLineaEntrada(linea)
		comandoAEjecutar := obtenerComando(entradaSeparada)
		parametros := entradaSeparada[COMANDO_A_EJECUTAR_POS_CMD+1:]

		switch comandoAEjecutar {
		case LOGIN:

			// Verificar parámetros de comando para ingresar votante.
			nombreUsuario, err := verificarParamsLogin(parametros)
			if err != nil {
				fmt.Println(err)
			}
			if !dicUsuario.Pertenece(nombreUsuario) {
				fmt.Println(errores.ErrorUsuarioNoExiste{})
				break
			}
			usuarioLogueado := dicUsuario.Obtener(nombreUsuario)
			nombre, err := usuarioLogueado.Login(dicUsuario, nombreUsuario, conectado)
			if err != nil {
				fmt.Println(err)
				return
			}
			//si todo sale bien este es el nuevo usuario activo
			*ptrUsuarioLogueado = dicUsuario.Obtener(nombreUsuario)
			conectado = true
			*ptrnombre = nombre
			fmt.Println(nombre)

		//votanteActual = vo.CrearVotante(dni)

		case LOGOUT:

			usuarioLogueado.Logout(dicUsuario, nombre)
			conectado = false
			fmt.Println("Adios")

		// case PUBLICAR:

		// case VER_SIGUIENTE_FEED:

		// case LIKEAR_POST:

		// case MOSTRAR_LIKES:

		default:
			err = errores.ErrorComandoDesconocido{}
		}
	}
}
func obtenerComando(campos []string) Comando {
	switch campos[COMANDO_A_EJECUTAR_POS_CMD] {
	case COMANDO_LOGIN:
		return LOGIN
	case COMANDO_LOGOUT:
		return LOGOUT
	case COMANDO_PUBLICAR:
		return PUBLICAR
	case COMANDO_VER_SIGUIENTE_FEED:
		return VER_SIGUIENTE_FEED
	case COMANDO_LIKEAR_POST:
		return LIKEAR_POST
	case COMANDO_MOSTRAR_LIKES:
		return MOSTRAR_LIKES
	}
	return ERROR
}

// SepararLineaEntrada realiza todas las acciones necesarias para separar la línea de comando en sus respectivos campos
// según el protocolo definido.
func separarLineaEntrada(linea string) []string {
	// Quitar espacios en el inicio y el fin si los hubiera.
	lineaSinSaltoDeLinea := strings.TrimSpace(linea)

	//Dividir la linea por cada espacio que contenga.
	partes := strings.Fields(lineaSinSaltoDeLinea)

	return partes
}
