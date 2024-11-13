package pila

const capPilaNueva = 2    // Capacidad a asignar al campo datos de una pilaDinamica nueva
const facRedimPila = 2    // Factor de redimension para el campo datos de una pilaDinamica
const multCantMinPila = 4 // Multiplicador para verificar la cantidad minima de una pilaDinamica para achicarla

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	inicial := pilaDinamica[T]{}
	inicial.datos = make([]T, capPilaNueva)

	return &inicial
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == cap(p.datos) {
		redimensionarPila(p, cap(p.datos)*facRedimPila+1)
	}

	p.datos[p.cantidad] = elem
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	tope := p.datos[p.cantidad-1]
	p.cantidad--

	if p.cantidad*multCantMinPila <= cap(p.datos) {
		redimensionarPila(p, cap(p.datos)/facRedimPila)
	}

	return tope
}

func redimensionarPila[T any](p *pilaDinamica[T], capNueva int) {
	nuevoDatos := make([]T, capNueva)
	copy(nuevoDatos, p.datos)
	p.datos = nuevoDatos
}
