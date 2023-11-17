package diccionario

import (
	TDAPila "tdas/pila"
)

const (
	_ZEROCMP          = 0
	_CANTIDAD_INICIAL = 0
)

type arbol[K comparable, V any] struct {
	izq, der *arbol[K, V]
	clave    K
	dato     V
	cant     int
	funcCmp  func(K, K) int
}

func CrearABB[K comparable, V any](funcCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	nuevoArbol := new(arbol[K, V])
	nuevoArbol.cant = _CANTIDAD_INICIAL
	nuevoArbol.funcCmp = funcCmp

	return nuevoArbol
}

// buscarNodo busca la unión entre el nodo padre
func (abb *arbol[K, V]) buscarNodo(clave K, unionPadreHijo ***arbol[K, V]) bool {
	if abb == nil {
		return false
	}

	if abb.funcCmp(abb.clave, clave) > _ZEROCMP {
		*unionPadreHijo = &abb.izq
		return abb.izq.buscarNodo(clave, unionPadreHijo)
	}
	if abb.funcCmp(abb.clave, clave) < _ZEROCMP {
		*unionPadreHijo = &abb.der
		return abb.der.buscarNodo(clave, unionPadreHijo)
	}

	return true
}

// modificarAlturas suma a la cantidad de nodos el valor pasado por argumento, exceptuando el nodo con la clave deseada.

func (abb *arbol[K, V]) modificarAlturas(clave K, valor int) {
	if abb == nil {
		return
	}

	if abb.funcCmp(abb.clave, clave) == _ZEROCMP {
		return
	}

	if abb.funcCmp(abb.clave, clave) > _ZEROCMP {
		abb.izq.modificarAlturas(clave, valor)
	} else {
		abb.der.modificarAlturas(clave, valor)
	}
	abb.cant += valor
}

// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (abb *arbol[K, V]) Guardar(clave K, dato V) {

	if abb.Cantidad() == _ZEROCMP {
		abb.cant++
		abb.clave = clave
		abb.dato = dato
		return
	}

	unionPadreHijo := &abb

	if abb.buscarNodo(clave, &unionPadreHijo) {
		(*unionPadreHijo).dato = dato
	} else {
		nuevoNodo := abb.crearNodo(clave, dato)
		(*unionPadreHijo) = nuevoNodo
		abb.modificarAlturas(clave, 1)
	}
}

// crearNodo crea e inicializa un nodo.
func (abb *arbol[K, V]) crearNodo(clave K, dato V) *arbol[K, V] {
	nuevoNodo := new(arbol[K, V])
	nuevoNodo.clave = clave
	nuevoNodo.dato = dato
	nuevoNodo.cant = 1
	nuevoNodo.funcCmp = abb.funcCmp

	return nuevoNodo
}

// Pertenece determina si una clave ya se encuentra en el diccionario, o no
func (abb *arbol[K, V]) Pertenece(clave K) bool {
	unionPadreHijo := &abb
	return abb.Cantidad() > _ZEROCMP && abb.buscarNodo(clave, &unionPadreHijo)
}

// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje 'La clave no pertenece al diccionario'
func (abb *arbol[K, V]) Obtener(clave K) V {
	unionPadreHijo := &abb

	if abb.Cantidad() == _ZEROCMP || !abb.buscarNodo(clave, &unionPadreHijo) {
		panic("La clave no pertenece al diccionario")
	}
	return (*unionPadreHijo).dato
}

// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
func (abb *arbol[K, V]) Borrar(clave K) V {
	unionPadreHijo := &abb

	if abb.Cantidad() == _ZEROCMP || !abb.buscarNodo(clave, &unionPadreHijo) {
		panic("La clave no pertenece al diccionario")
	}

	if *unionPadreHijo == abb {
		return abb.borrarRaiz()
	}

	abb.modificarAlturas(clave, -1)

	return abb.borrarNodo(unionPadreHijo)
}

// reemplazarDatos reemplaza la clave y el dato del nodo que invoca la primitiva por los almacenados del nodo pasado por argumento.
func (abb *arbol[K, V]) reemplazarDatos(nodoReemplazante *arbol[K, V]) {
	abb.clave = nodoReemplazante.clave
	abb.dato = nodoReemplazante.dato
}

// reemplazarHijos reemplaza los hijos del nodo que invoca la función por los del nodo pasado por argumento.
func (abb *arbol[K, V]) reemplazarHijos(nodoReemplazante *arbol[K, V]) {
	abb.izq = nodoReemplazante.izq
	abb.der = nodoReemplazante.der
}

// Funciones de borrado de la raiz.
func (abb *arbol[K, V]) borrarRaiz() V {
	if abb.Cantidad() == 1 {
		abb.cant = _CANTIDAD_INICIAL
		return abb.dato
	}

	datoBorrado := abb.dato

	if abb.izq != nil && abb.der == nil {
		abb.reemplazarDatos(abb.izq)
		abb.reemplazarHijos(abb.izq)
		abb.cant--
	} else if abb.izq == nil && abb.der != nil {
		abb.reemplazarDatos(abb.der)
		abb.reemplazarHijos(abb.der)
		abb.cant--
	} else {
		reemplazante := abb.der.buscarReemplazante(&abb.der)

		abb.modificarAlturas((*reemplazante).clave, -1)

		abb.reemplazarDatos((*reemplazante))

		abb.borrarNodo(reemplazante)
	}
	return datoBorrado
}

// Funciones de borrado de nodos distintos a la raiz.
func (abb *arbol[K, V]) borrarNodo(unionPadreHijo **arbol[K, V]) V {
	if (*unionPadreHijo).izq == nil && (*unionPadreHijo).der == nil {
		return abb.borrarHoja(unionPadreHijo)
	} else if (*unionPadreHijo).izq != nil && (*unionPadreHijo).der != nil {
		return abb.borrarNodoDosHijos(unionPadreHijo)
	}

	return abb.borrarNodoUnHijo(unionPadreHijo)
}

// borrarHoja borra la hoja almacenada en la posición de memoria pasada por parámetro y devuelve el dato almacenado en ella.
func (abb *arbol[K, V]) borrarHoja(unionPadreHijo **arbol[K, V]) V {
	datoBorrado := (*unionPadreHijo).dato

	(*unionPadreHijo) = nil

	return datoBorrado
}

// borrarNodoUnHijoDer borra un nodo con un único hijo, que es el nodo almacenado en la posición de memoria pasada por parámetro,
// y devuelve el dato almacenado en él.
func (abb *arbol[K, V]) borrarNodoUnHijo(unionPadreHijo **arbol[K, V]) V {
	datoBorrado := (*unionPadreHijo).dato

	if (*unionPadreHijo).izq != nil && (*unionPadreHijo).der == nil {
		(*unionPadreHijo) = (*unionPadreHijo).izq
	} else {
		(*unionPadreHijo) = (*unionPadreHijo).der
	}

	return datoBorrado
}

// borrarNodoDosHijos borra un nodo con dos hijos, que es el nodo almacenado en la posición de memoria pasada por parámetro,
// y devuelve el dato almacenado en él.
func (abb *arbol[K, V]) borrarNodoDosHijos(unionPadreHijo **arbol[K, V]) V {

	// Guarda el dato del nodo a borrar.
	datoBorrado := (*unionPadreHijo).dato

	// Se busca un nodo reemplazante al nodo a borrar.
	unionReemplazante := (*unionPadreHijo).der.buscarReemplazante(&(*unionPadreHijo).der)

	// Copia el dato y la clave del mejor reemplazante en el nodo a sustituir.
	(*unionPadreHijo).clave = (*unionReemplazante).clave
	(*unionPadreHijo).dato = (*unionReemplazante).dato

	// Borra el nodo mejor reemplazante.
	if (*unionReemplazante).izq == nil && (*unionReemplazante).der == nil {
		abb.borrarHoja(unionReemplazante)
	} else {
		abb.borrarNodoUnHijo(unionReemplazante)
	}

	return datoBorrado
}

// buscarReemplazante busca el nodo que reempĺazará al
func (abb *arbol[K, V]) buscarReemplazante(candidato **arbol[K, V]) **arbol[K, V] {
	if abb.izq == nil {
		return candidato
	}

	return abb.izq.buscarReemplazante(&abb.izq)
}

// Cantidad devuelve la cantidad de elementos dentro del diccionario
func (abb *arbol[K, V]) Cantidad() int {
	return abb.cant
}

// Iterar itera internamente el diccionario, aplicando la función pasada por parámetro a todos los elementos del mismo
func (abb *arbol[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	if abb.Cantidad() == _CANTIDAD_INICIAL {
		return
	}
	abb.iterarRangoRec(nil, nil, visitar)

}

// IterarRango itera sólo incluyendo a los elementos que se encuentren comprendidos en el rango indicado, incluyéndolos en caso de encontrarse
func (abb *arbol[K, V]) IterarRango(desde, hasta *K, visitar func(clave K, dato V) bool) {
	if abb.Cantidad() == _CANTIDAD_INICIAL {
		return
	}
	abb.iterarRangoRec(desde, hasta, visitar)
}

func (abb *arbol[K, V]) iterarRangoRec(desde, hasta *K, visitar func(clave K, dato V) bool) bool {
	if abb == nil {
		return true
	}

	if desde != nil && abb.funcCmp(abb.clave, *desde) < _ZEROCMP {
		return abb.der.iterarRangoRec(desde, hasta, visitar)
	}

	if hasta != nil && abb.funcCmp(abb.clave, *hasta) > _ZEROCMP {
		return abb.izq.iterarRangoRec(desde, hasta, visitar)
	}

	if !abb.izq.iterarRangoRec(desde, hasta, visitar) || !visitar(abb.clave, abb.dato) {
		return false
	}
	return abb.der.iterarRangoRec(desde, hasta, visitar)
}

// ----------------------------------------
// ------------- iter externo-------------
// ----------------------------------------

type iterAbb[K comparable, V any] struct {
	pila  TDAPila.Pila[*arbol[K, V]]
	desde *K
	hasta *K
	abb   *arbol[K, V]
}

// Iterador devuelve un IterDiccionario para este Diccionario
func (abb *arbol[K, V]) Iterador() IterDiccionario[K, V] {
	iter := abb.IteradorRango(nil, nil)
	return iter
}

// Iterador Rango crea un IterDiccionario que sólo itere por las claves que se encuentren en el rango indicado
func (abb *arbol[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	iteradorRango := new(iterAbb[K, V])

	iteradorRango.pila = TDAPila.CrearPilaDinamica[*arbol[K, V]]()
	iteradorRango.desde = desde
	iteradorRango.hasta = hasta
	iteradorRango.abb = abb

	if abb.cant == _CANTIDAD_INICIAL {
		return iteradorRango
	}
	iteradorRango.apilarTodoIzqRango(iteradorRango.abb)
	return iteradorRango
}

// Apilar todos los nodos hacia la izquierda si esta en el rango
func (iter *iterAbb[K, V]) apilarTodoIzqRango(abb *arbol[K, V]) {
	if abb == nil {
		return
	}

	if iter.desde != nil && iter.abb.funcCmp(abb.clave, *iter.desde) < _ZEROCMP {
		iter.apilarTodoIzqRango(abb.der)
		return
	}
	if iter.hasta != nil && iter.abb.funcCmp(abb.clave, *iter.hasta) > _ZEROCMP {
		iter.apilarTodoIzqRango(abb.izq)
		return
	}
	iter.pila.Apilar(abb)
	iter.apilarTodoIzqRango(abb.izq)

}

// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado el iterador hay un elemento.
func (iter iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()

}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador. Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
func (iter iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

// Siguiente si HaySiguiente, devuelve la clave actual (equivalente a VerActual, pero únicamente la clave), y además avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe entrar en pánico con mensaje 'El iterador termino de iterar'
func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodo := iter.pila.Desapilar()
	if nodo.der != nil {
		iter.apilarTodoIzqRango(nodo.der)
	}

}
