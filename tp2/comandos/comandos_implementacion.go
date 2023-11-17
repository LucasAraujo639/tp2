package comandos

import (
	"tp2/errores"
)

// Cantidad de elementos de cada comando al invocar:
const (
	CANT_PARAMETROS_COMANDO_LOGIN              = 1
	CANT_PARAMETROS_COMANDO_LOGOUT             = 0
	CANT_PARAMETROS_COMANDO_VER_SIGUIENTE_FEED = 0
	CANT_PARAMETROS_COMANDO_LIKEAR_POST        = 1
	CANT_PARAMETROS_COMANDO_MOSTRAR_LIKES      = 1
)

// ejemplos comandos
// login chicho1994
// logout
// publicar Tiene todo el dinero del mundo...
// ver_siguiente_feed
// likear_post 0
// mostrar_likes 0

// Posición de comando a ejecutar en entrada:
const COMANDO_A_EJECUTAR_POS_CMD int = 0

// -------------------------------------------------------------------------------------------------------------
//                                 Comando de ingreso de votante nuevo.
// -------------------------------------------------------------------------------------------------------------

const (
	COMANDO_POS_CMD = 0
)

// verificarParamsLogin comprueba que los parámetros pasados por argumento sean validos.
func verificarParamsLogin(parametros []string) (string, error) {
	if len(parametros) != CANT_PARAMETROS_COMANDO_LOGIN {
		return "", errores.ErrorMalaInvocacionComando{}
	}

	usuario := parametros[COMANDO_POS_CMD]

	return usuario, nil
}
