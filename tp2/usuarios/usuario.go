package usuarios

import TDAHash "tdas/diccionario"

type Usuario interface {
	// Login perimte al usuario loggearse
	Login(TDAHash.Diccionario[string, Usuario], string, bool) (string, error)

	//Logout permite al usuario desloguearse
	Logout(TDAHash.Diccionario[string, Usuario], string) error
	// VerNombre devuelve el nombre del usuario.
	VerNombre() string
	VerEstado() Estado
	SetEstado(Estado)

	// // CrearPost crea un post con el texto y el ID pasados por argumento.
	// CrearPost(int, string) TDAPost.Post

	// // VerSiguientePost devuelve el siguiente post a ver en el feed.
	// VerSiguientePost() (TDAPost.Post, error)

	// // LikearPost le da "Me gusta" al post con el ID indicado.
	// LikearPost(int) error

	// // GuardarPostFeed guarda en el feed del usuario un post según la afinidad definida.
	// GuardarPostFeed(TDAPost.Post, int)
}
