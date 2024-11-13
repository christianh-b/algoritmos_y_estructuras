package pila

const (
	capacidadInicial = 5
	cantidadInicial  = 0
	k                = 2
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		make([]T, capacidadInicial),
		cantidadInicial,
	}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == cantidadInicial
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		pila.mensajePanic()
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	capacidadActual := cap(pila.datos)
	if pila.cantidad == capacidadActual {
		dobleDeCapacidad := capacidadActual * k
		pila.redimensionar(dobleDeCapacidad)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		pila.mensajePanic()
	}
	elementoADesapilar := pila.VerTope()
	pila.cantidad--
	capacidadActual := cap(pila.datos)
	cantidadCuadriplicada := pila.cantidad * k * k
	if capacidadActual >= cantidadCuadriplicada && capacidadActual > capacidadInicial {
		mitadDeCapacidad := capacidadActual / k
		pila.redimensionar(mitadDeCapacidad)
	}
	return elementoADesapilar
}

func (pila *pilaDinamica[T]) mensajePanic() {
	panic("La pila esta vacia")
}

func (pila *pilaDinamica[T]) redimensionar(tamaño int) {
	if tamaño < capacidadInicial {
		tamaño = capacidadInicial
	}
	pilaRedimensionada := make([]T, tamaño)
	copy(pilaRedimensionada, pila.datos[:pila.cantidad])
	pila.datos = pilaRedimensionada
}
