package errores

type ErrorUsuarioNoLoggeado struct{}

func (e ErrorUsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorUsuarioNoExiste struct{}

func (e ErrorUsuarioNoExiste) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioYaLoggeado struct{}

func (e ErrorUsuarioYaLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorPostInexistente struct{}

func (e ErrorPostInexistente) Error() string {
	return "Error: Post inexistente"
}

type ErrorPostSinLikes struct{}

func (e ErrorPostSinLikes) Error() string {
	return "Error: Post sin likes"
}

type ErrorNoMasPosts struct{}

func (e ErrorNoMasPosts) Error() string {
	return "Error: no hay mas posts para ver"
}

type ErrorLecturaArchivo struct{}

func (e ErrorLecturaArchivo) Error() string {
	return "Error: no se pudo leer el archivo"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "Error: Faltan Parametros"
}

type ErrorMalaInvocacionComando struct{}

func (e ErrorMalaInvocacionComando) Error() string {
	return "ERROR: Comando mal invocado"
}

type ErrorComandoDesconocido struct{}

func (e ErrorComandoDesconocido) Error() string {
	return "ERROR: Comando desconocido"
}
