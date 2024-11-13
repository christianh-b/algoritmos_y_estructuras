package diccionario

import "fmt"

type estadoCeldaHash int

const (
	VACIO estadoCeldaHash = iota
	BORRADO
	OCUPADO
)

const (
	M     = 0x5bd1e995
	BIG_M = 0xc6a4a7935bd1e995
	R     = 24
	BIG_R = 47

	CAPACIDAD_INICIAL     = 11
	FACTOR_REDIMENSION    = 2
	MAX_FACTOR_CARGA      = 0.7
	MIN_FACTOR_CARGA_OCUP = 0.2
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoCeldaHash
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tamano   int
	borrados int
}

func crearTablaHashCerrado[K comparable, V any](tam int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], tam)
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{
		tabla:  crearTablaHashCerrado[K, V](CAPACIDAD_INICIAL),
		tamano: CAPACIDAD_INICIAL,
	}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	pos := hash.buscarClave(clave)

	if pos != -1 {
		hash.tabla[pos].dato = dato
	} else {
		if hash.hayQueAgrandar() {
			hash.redimensionar(hash.tamano * FACTOR_REDIMENSION)
		}

		pos = hash.obtenerPosicion(clave)
		for hash.tabla[pos].estado == OCUPADO {
			pos = (pos + 1) % hash.tamano
		}

		if hash.tabla[pos].estado == BORRADO {
			hash.borrados--
		}
		hash.tabla[pos].clave = clave
		hash.tabla[pos].dato = dato
		hash.tabla[pos].estado = OCUPADO
		hash.cantidad++
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	return hash.buscarClave(clave) != -1
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos := hash.buscarClave(clave)

	hash.validarPosClave(pos)

	return hash.tabla[pos].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	pos := hash.buscarClave(clave)

	hash.validarPosClave(pos)

	datoABorrar := hash.tabla[pos].dato
	hash.tabla[pos].estado = BORRADO
	hash.borrados++
	hash.cantidad--

	if hash.hayQueAchicar() {
		hash.redimensionar(hash.tamano / FACTOR_REDIMENSION)
	}

	return datoABorrar
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			seguir := visitar(celda.clave, celda.dato)
			if !seguir {
				break
			}
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iterHashCerrado[K, V]{
		posActual: -1,
		tabla:     hash.tabla,
		tamHash:   hash.tamano,
	}
	iterador.posActual = iterador.buscarSigPosOcup()

	return iterador
}

func (hash *hashCerrado[K, V]) buscarClave(clave K) int {
	pos := hash.obtenerPosicion(clave)

	for hash.tabla[pos].estado != VACIO {
		if hash.tabla[pos].estado == OCUPADO && hash.tabla[pos].clave == clave {
			return pos
		}
		pos = (pos + 1) % hash.tamano
	}

	return -1
}

func (hash *hashCerrado[K, V]) obtenerPosicion(clave K) int {
	claveHasheada := murmurHash2(convertirABytes(clave), 0)

	return int(claveHasheada) % hash.tamano
}

func (hash *hashCerrado[K, V]) validarPosClave(posicion int) {
	if posicion == -1 {
		panic("La clave no pertenece al diccionario")
	}
}

func (hash *hashCerrado[K, V]) hayQueAgrandar() bool {
	factorDeCarga := float64(hash.cantidad+hash.borrados) / float64(hash.tamano)

	return factorDeCarga >= MAX_FACTOR_CARGA
}

func (hash *hashCerrado[K, V]) hayQueAchicar() bool {
	factorDeCargaOcup := float64(hash.cantidad) / float64(hash.tamano)

	return factorDeCargaOcup <= MIN_FACTOR_CARGA_OCUP && hash.cantidad > CAPACIDAD_INICIAL
}

func (hash *hashCerrado[K, V]) redimensionar(tamNuevo int) {
	tablaRedimensionada := crearTablaHashCerrado[K, V](tamNuevo)
	hash.tamano = tamNuevo

	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			pos := hash.obtenerPosicion(celda.clave)
			for tablaRedimensionada[pos].estado == OCUPADO {
				pos = (pos + 1) % tamNuevo
			}
			tablaRedimensionada[pos] = celda
		}
	}

	hash.tabla = tablaRedimensionada
	hash.borrados = 0
}

func murmurHash2(data []byte, seed uint32) (h uint32) {
	var k uint32
	h = seed ^ uint32(len(data))

	for l := len(data); l >= 4; l -= 4 {
		k = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
		h, k = mmix(h, k)
		data = data[4:]
	}

	switch len(data) {
	case 3:
		h ^= uint32(data[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(data[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(data[0])
		h *= M
	}

	h ^= h >> 13
	h *= M
	h ^= h >> 15

	return
}

func mmix(h uint32, k uint32) (uint32, uint32) {
	k *= M
	k ^= k >> R
	k *= M
	h *= M
	h ^= k

	return h, k
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

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
