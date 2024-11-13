package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

func crearNodoLista[T any](elem T, proximo *nodoLista[T]) nodoLista[T] {
	return nodoLista[T]{dato: elem, prox: proximo}
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := crearNodoLista(elem, l.primero)

	if l.EstaVacia() {
		l.ultimo = &nuevoNodo
	}

	l.primero = &nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevoNodo := crearNodoLista(elem, nil)

	if l.EstaVacia() {
		l.primero = &nuevoNodo
	} else {
		l.ultimo.prox = &nuevoNodo
	}

	l.ultimo = &nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	l.validarLista()

	primerDato := l.primero.dato
	l.primero = l.primero.prox
	l.largo--

	if l.EstaVacia() {
		l.ultimo = nil
	}

	return primerDato
}

func (l *listaEnlazada[T]) VerPrimero() T {
	l.validarLista()

	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	l.validarLista()

	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero

	for actual != nil {
		seguir := visitar(actual.dato)

		if !seguir {
			break
		}
		actual = actual.prox
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	it := new(iterListaEnlazada[T])
	it.l = l
	it.actual = it.l.primero

	return it
}

func (l *listaEnlazada[T]) validarLista() {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
}
