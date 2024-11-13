package lista

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	l        *listaEnlazada[T]
}

func (it *iterListaEnlazada[T]) VerActual() T {
	it.validarIterador()

	return it.actual.dato
}

func (it *iterListaEnlazada[T]) HaySiguiente() bool {
	return it.actual != nil
}

func (it *iterListaEnlazada[T]) Siguiente() {
	it.validarIterador()

	it.anterior = it.actual
	it.actual = it.actual.prox
}

func (it *iterListaEnlazada[T]) Insertar(elem T) {
	nodo := crearNodoLista(elem, nil)

	if it.anterior == nil {
		if it.actual == nil {
			it.l.primero = &nodo
			it.l.ultimo = &nodo
			it.actual = it.l.primero
		} else {
			nodo.prox = it.actual
			it.actual = &nodo
			it.l.primero = it.actual
		}
	} else {
		it.anterior.prox = &nodo
		nodo.prox = it.actual
		it.actual = &nodo

		if it.actual.prox == nil {
			it.l.ultimo = it.actual
		}
	}
	it.l.largo++
}

func (it *iterListaEnlazada[T]) Borrar() T {
	it.validarIterador()

	elementoaBorrar := it.actual.dato

	if it.anterior == nil {
		it.l.primero = it.actual.prox
		it.actual = it.actual.prox
	} else {
		it.anterior.prox = it.actual.prox

		if it.actual.prox == nil {
			it.l.ultimo = it.anterior
		}
		it.actual = it.anterior.prox
	}
	it.l.largo--

	return elementoaBorrar
}

func (it *iterListaEnlazada[T]) validarIterador() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
