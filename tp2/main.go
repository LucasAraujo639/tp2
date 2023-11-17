package main

import (
	"fmt"
	"tp2/comandos"
	"tp2/datos"
)

func main() {
	// Carga de datos...
	// Obtener datos de invocacion.
	//params := datos.ObtenerParametrosEjecucion()
	// err := datos.VerificarParametrosEjecucion(params)
	// if err != nil {
	// 	fmt.Fprintf(os.Stdout, "%s\n", err.Error())
	// 	return
	// }
	// archivoUsuarios, err := datos.AbrirArchivo(datos.ObtenerNombreArchivoUsuarios(params))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	archivoUsuarios, err := datos.AbrirArchivotest()
	if err != nil {
		fmt.Println(err)
		return
	}

	//se inicializa el hash con los usuarios adentro
	hashUsuario, err := datos.CargarUsuarios(archivoUsuarios)
	if err != nil {
		fmt.Println(err)
		return
	}
	comandos.LectorComandos(hashUsuario)

}
