package cola_prioridad

const (
	POSICION_INICIAL       = 0
	CAPACIDAD_INICIAL      = 11
	FACTOR_REDIMENSION     = 2
	MAX_FACTOR_REDIMENSION = 0.7
	MIN_FACTOR_REDIMENSION = 0.25
)

type heap[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, CAPACIDAD_INICIAL)
	heap.cmp = funcion_cmp

	return heap
}

func CrearHeapArr[T any](elementos []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, len(elementos))
	heap.cant = len(elementos)
	heap.cmp = funcion_cmp

	copy(heap.datos, elementos)
	heapify(heap.datos, heap.cmp)

	return heap
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.hayQueAgrandar() {
		heap.redimensionar(len(heap.datos) * FACTOR_REDIMENSION)
	}

	heap.datos[heap.cant] = elem
	upHeap(heap.datos, heap.cant, heap.cmp)
	heap.cant++
}

func (heap *heap[T]) VerMax() T {
	heap.validarHeap()

	return heap.datos[POSICION_INICIAL]
}

func (heap *heap[T]) Desencolar() T {
	heap.validarHeap()
	elemBorrado := heap.VerMax()

	swap(&heap.datos[POSICION_INICIAL], &heap.datos[heap.cant-1])
	heap.cant--
	downHeap(heap.datos, POSICION_INICIAL, heap.cant-1, heap.cmp)

	if heap.hayQueAchicar() {
		heap.redimensionar(len(heap.datos) / FACTOR_REDIMENSION)
	}

	return elemBorrado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cant
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)

	for fin := len(elementos) - 1; fin > 0; fin-- {
		swap(&elementos[POSICION_INICIAL], &elementos[fin])
		downHeap(elementos, POSICION_INICIAL, fin-1, funcion_cmp)
	}
}

func (heap *heap[T]) validarHeap() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (heap *heap[T]) hayQueAgrandar() bool {
	factorRedimension := float64(heap.cant) / float64(len(heap.datos))

	return factorRedimension >= MAX_FACTOR_REDIMENSION || len(heap.datos) < CAPACIDAD_INICIAL
}

func (heap *heap[T]) hayQueAchicar() bool {
	factorRedimension := float64(heap.cant) / float64(len(heap.datos))

	return factorRedimension <= MIN_FACTOR_REDIMENSION && len(heap.datos) > CAPACIDAD_INICIAL
}

func (heap *heap[T]) redimensionar(tamano int) {
	if tamano < CAPACIDAD_INICIAL {
		tamano = CAPACIDAD_INICIAL
	}

	heapRedimensionado := make([]T, tamano)
	copy(heapRedimensionado, heap.datos[:heap.cant])
	heap.datos = heapRedimensionado
}

func swap[T any](elem1 *T, elem2 *T) {
	*elem1, *elem2 = *elem2, *elem1
}

func cumpleCondicionDeHeap[T any](elementos []T, posPadre int, posHijo int, funcion_cmp func(T, T) int) bool {
	return funcion_cmp(elementos[posPadre], elementos[posHijo]) > 0
}

func upHeap[T any](elementos []T, posAct int, funcion_cmp func(T, T) int) {
	posPadre := (posAct - 1) / 2

	if posAct == POSICION_INICIAL || cumpleCondicionDeHeap(elementos, posPadre, posAct, funcion_cmp) {
		return
	}

	swap(&elementos[posAct], &elementos[posPadre])
	upHeap(elementos, posPadre, funcion_cmp)
}

func downHeap[T any](elementos []T, ini int, fin int, funcion_cmp func(T, T) int) {
	padre := ini
	hijoIzq := (2 * ini) + 1
	hijoDer := (2 * ini) + 2

	if ini >= fin || hijoIzq > fin && hijoDer > fin {
		return
	}

	if hijoIzq <= fin && !cumpleCondicionDeHeap(elementos, padre, hijoIzq, funcion_cmp) {
		padre = hijoIzq
	}

	if hijoDer <= fin && !cumpleCondicionDeHeap(elementos, padre, hijoDer, funcion_cmp) {
		padre = hijoDer
	}

	if padre == ini {
		return
	}

	swap(&elementos[ini], &elementos[padre])
	downHeap(elementos, padre, fin, funcion_cmp)
}

func heapify[T any](elementos []T, funcion_cmp func(T, T) int) {
	for i := (len(elementos) - 1) / 2; i >= 0; i-- {
		downHeap(elementos, i, len(elementos)-1, funcion_cmp)
	}
}
