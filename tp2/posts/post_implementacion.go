package posts

import (
	"strings"
	TDADicccionario "tdas/diccionario"
)

type postImplementacion struct {
	id    int
	creador string
	texto string
	likes TDADicccionario.DiccionarioOrdenado[string, int]
}

// Crea el post.
func CrearPost(id int, usr string, texto string) Post {
	postNuevo := new(postImplementacion)

	postNuevo.creador = usr
	postNuevo.id = id
	postNuevo.texto = texto
	postNuevo.likes = TDADicccionario.CrearABB[string, int](func(s1, s2 string) int { return strings.Compare(s1,s2) })

	return postNuevo
}

// VerID devuelve el ID del post.
func (post *postImplementacion) VerID() int {
	return post.id
}

// VerPost devuelve el texto de la publicación.
func (post *postImplementacion) VerPost() string {
	return post.texto
}

// Likear le da un "Me gusta" al post. Se pasa por parámetro el nombre del usuario que le da "Me gusta" al post.
// No se verifica la existencia del usuario pasado por argumento.
func (post *postImplementacion) Likear(usr string) {
	post.likes.Guardar(usr, 0)
}

// VerLikes devuelve los nombres de los usuarios que dieron "Me gusta" al post ordenados en orden alfabético.
func (post *postImplementacion) VerLikes() []string {
	usuariosLikes := make([]string, post.likes.Cantidad())
	indice := 0

	post.likes.Iterar(func(usr string, _ int) bool {
		usuariosLikes[indice] = usr
		return true
	})

	return usuariosLikes
}
