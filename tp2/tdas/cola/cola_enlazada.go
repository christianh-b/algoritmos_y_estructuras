package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{
		dato: dato,
	}
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	cola.validarCola()
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nuevoNodo := nodoCrear(elemento)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
	}
	cola.ultimo = nuevoNodo

}

func (cola *colaEnlazada[T]) Desencolar() T {
	cola.validarCola()
	elemento := cola.VerPrimero()
	cola.primero = cola.primero.prox
	return elemento
}

func (cola *colaEnlazada[T]) validarCola() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}
