package diccionario

/*
type iterHashCerrado[K comparable, V any] struct {
	posActual int
	tabla     []celdaHash[K, V]
	tamHash   int
}

func (it *iterHashCerrado[K, V]) HaySiguiente() bool {
	return it.posActual != -1
}

func (it *iterHashCerrado[K, V]) VerActual() (K, V) {
	it.validarIterador()

	return it.tabla[it.posActual].clave, it.tabla[it.posActual].dato
}

func (it *iterHashCerrado[K, V]) Siguiente() {
	it.validarIterador()

	it.posActual = it.buscarSigPosOcup()
}

func (it *iterHashCerrado[K, V]) buscarSigPosOcup() int {
	for i := it.posActual + 1; i < it.tamHash; i++ {
		if it.tabla[i].estado == OCUPADO {
			return i
		}
	}

	return -1
}

func (it *iterHashCerrado[K, V]) validarIterador() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
*/
