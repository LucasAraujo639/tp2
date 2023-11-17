package datos

import (
	"bufio"
	"os"
	"strings"
	TDAHash "tdas/diccionario"
	"tp2/errores"
	"tp2/usuarios"
)

// AbrirArchivo intenta abrir el archivo con nombre pasado como argumento y devuelve un error según el criterio usitlizado.
func AbrirArchivo(nombreArchivo string) (*os.File, error) {
	archivoAbierto, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, errores.ErrorLecturaArchivo{}
	}
	return archivoAbierto, nil
}
func AbrirArchivotest() (*os.File, error) {
	archivoAbierto, err := os.Open("usuarios.txt")
	if err != nil {
		return nil, errores.ErrorLecturaArchivo{}
	}
	return archivoAbierto, nil
}

func CargarUsuarios(archivo *os.File) (TDAHash.Diccionario[string, usuarios.Usuario], error) {
	defer archivo.Close()
	dic := TDAHash.CrearHash[string, usuarios.Usuario]()
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		nombreUsuario := strings.TrimSpace(scanner.Text())
		if nombreUsuario == "" {
			continue
		} else {
			nuevoUsuario := usuarios.CrearUsuario(nombreUsuario)
			dic.Guardar(nombreUsuario, nuevoUsuario) // guardo en los dos nombre de uusario porque nose que guardar por ahor
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, errores.ErrorLecturaArchivo{}
	}

	return dic, nil
}
