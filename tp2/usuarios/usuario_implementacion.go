package usuarios

import (
	TDAColaPrioridad "tdas/cola_prioridad"
	TDAHash "tdas/diccionario"
	"tp2/errores"
	TDAPost "tp2/posts"
)

type Estado int

const (
	DESCONECTADO Estado = iota
	CONECTADO
)

type usuarioImplementacion struct {
	nombre string
	feed   TDAColaPrioridad.ColaPrioridad[TDAPost.Post]
	estado Estado
}

func CrearUsuario(nombre string) Usuario {
	nuevoUsuario := new(usuarioImplementacion)
	nuevoUsuario.nombre = nombre
	nuevoUsuario.estado = DESCONECTADO
	return nuevoUsuario
}
func (usr *usuarioImplementacion) Login(dic TDAHash.Diccionario[string, Usuario], usuario string, conectado bool) (string, error) {
	if !dic.Pertenece(usuario) {
		return "", errores.ErrorUsuarioNoExiste{}
	}
	if conectado {
		return "", errores.ErrorUsuarioYaLoggeado{}
	}
	usuarioObtenido := dic.Obtener(usuario)
	estado := usuarioObtenido.VerEstado()
	if estado == CONECTADO {
		return "", errores.ErrorUsuarioYaLoggeado{}
	}
	usuarioObtenido.SetEstado(CONECTADO)
	return ("Hola " + usuario), nil

}

func (usr *usuarioImplementacion) Logout(dic TDAHash.Diccionario[string, Usuario], usuario string) error {
	if !dic.Pertenece(usuario) {
		return errores.ErrorUsuarioNoExiste{}
	}
	usuarioObtenido := dic.Obtener(usuario)
	estado := usuarioObtenido.VerEstado()
	if estado == DESCONECTADO {
		return errores.ErrorUsuarioNoLoggeado{}
	}
	usuarioObtenido.SetEstado(DESCONECTADO)
	return nil
}

// VerNombre devuelve el nombre del usuario.
func (usr *usuarioImplementacion) VerNombre() string {
	return usr.nombre
}
func (usr *usuarioImplementacion) VerEstado() Estado {
	return usr.estado
}
func (usr *usuarioImplementacion) SetEstado(nuevoEstado Estado) {
	usr.estado = nuevoEstado

}

// CrearPost crea un post con el texto y el ID pasados por argumento.
func (usr *usuarioImplementacion) CrearPost(id int, texto string) TDAPost.Post {
	return TDAPost.CrearPost(id, usr.VerNombre(), texto)
}

// VerSiguientePost devuelve el siguiente post a ver en el feed.
func (usr *usuarioImplementacion) VerSiguientePost() (TDAPost.Post, error) {
	if usr.feed.EstaVacia() {
		return nil, errores.ErrorNoMasPosts{}
	}

	return usr.feed.Desencolar(), nil
}

// LikearPost le da "Me gusta" al post indicado.
func (usr *usuarioImplementacion) LikearPost(TDAPost.Post) error {
	return nil
}

// GuardarPostFeed guarda en el feed del usuario un post.
func (usr *usuarioImplementacion) GuardarPostFeed(TDAPost.Post) {

}
