package diccionario

import TDAPila "tdas/pila"

type nodoABB[K comparable, V any] struct {
	clave K
	dato  V
	izq   *nodoABB[K, V]
	der   *nodoABB[K, V]
}

func crearNodoABB[K comparable, V any](clave K, dato V) nodoABB[K, V] {
	return nodoABB[K, V]{
		clave: clave,
		dato:  dato,
	}
}

type abb[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		cmp: funcion_cmp,
	}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	refARefABuscado := abb.obtenerRefARefANodo(clave)

	if (*refARefABuscado) != nil {
		(*refARefABuscado).dato = dato
	} else {
		nodoNuevo := crearNodoABB(clave, dato) // creo una hoja
		(*refARefABuscado) = &nodoNuevo        // agrego hoja al abb
		abb.cantidad++
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	refARefABuscado := abb.obtenerRefARefANodo(clave)

	return (*refARefABuscado) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	refARefABuscado := abb.obtenerRefARefANodo(clave)

	abb.validarRefAClave(*refARefABuscado)

	return (*refARefABuscado).dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	refARefABuscado := abb.obtenerRefARefANodo(clave)

	abb.validarRefAClave(*refARefABuscado)

	datoABorrar := (*refARefABuscado).dato

	if (*refARefABuscado).izq != nil && (*refARefABuscado).der != nil {
		refARefAReemplazo := (*refARefABuscado).obtenerRefARefAReemplazo()

		(*refARefABuscado).clave = (*refARefAReemplazo).clave
		(*refARefABuscado).dato = (*refARefAReemplazo).dato

		eliminarNodo(refARefAReemplazo)

	} else {
		eliminarNodo(refARefABuscado)
	}
	abb.cantidad--

	return datoABorrar
}

func (nodo *nodoABB[K, V]) obtenerRefARefAReemplazo() **nodoABB[K, V] {
	if nodo.der.izq == nil {
		return &nodo.der
	}

	return nodo.der.obtenerRefARefAMinimo()
}

func (nodo *nodoABB[K, V]) obtenerRefARefAMinimo() **nodoABB[K, V] {
	if nodo.izq.izq == nil {
		return &nodo.izq
	}

	return nodo.izq.obtenerRefARefAMinimo()
}

func eliminarNodo[K comparable, V any](refARefANodo **nodoABB[K, V]) {
	if (*refARefANodo).izq != nil {
		(*refARefANodo) = (*refARefANodo).izq
	} else {
		(*refARefANodo) = (*refARefANodo).der
	}
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(K, V) bool) {
	abb.raiz.iterarRango(nil, nil, abb.cmp, visitar)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iterABB[K, V]{
		cmp:             abb.cmp,
		desde:           nil,
		hasta:           nil,
		pilaIteraciones: TDAPila.CrearPilaDinamica[*nodoABB[K, V]](),
	}
	iterador.apilarIteraciones(abb.raiz)

	return iterador
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	abb.raiz.iterarRango(desde, hasta, abb.cmp, visitar)
}

func (nodo *nodoABB[K, V]) iterarRango(desde *K, hasta *K, cmp func(K, K) int, visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}
	seguir := true

	if desde == nil || cmp(*desde, nodo.clave) < 0 {
		seguir = nodo.izq.iterarRango(desde, hasta, cmp, visitar)
	}
	if seguir && nodo.estaEnRango(desde, hasta, cmp) {
		seguir = visitar(nodo.clave, nodo.dato)
	}
	if seguir && (hasta == nil || cmp(nodo.clave, *hasta) < 0) {
		seguir = nodo.der.iterarRango(desde, hasta, cmp, visitar)
	}

	return seguir
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := &iterABB[K, V]{
		cmp:             abb.cmp,
		desde:           desde,
		hasta:           hasta,
		pilaIteraciones: TDAPila.CrearPilaDinamica[*nodoABB[K, V]](),
	}
	iterador.apilarIteraciones(abb.raiz)

	return iterador
}

// Devuelve una ref al campo (raiz de abb o hijo de nodo) que contiene una ref al nodo buscado
func (abb *abb[K, V]) obtenerRefARefANodo(claveBuscada K) **nodoABB[K, V] {
	if abb.raiz == nil {
		return &abb.raiz
	} else {
		if abb.cmp(abb.raiz.clave, claveBuscada) == 0 {
			return &abb.raiz
		} else {
			return abb.raiz.obtenerRefARefAHijo(claveBuscada, abb.cmp)
		}
	}
}

// Devuelve una ref al campo (hijo de nodo) que contiene una ref al nodo buscado
func (nodo *nodoABB[K, V]) obtenerRefARefAHijo(claveBuscada K, cmp func(K, K) int) **nodoABB[K, V] {
	if nodo.izq != nil && cmp(nodo.izq.clave, claveBuscada) == 0 {
		return &nodo.izq
	}
	if nodo.der != nil && cmp(nodo.der.clave, claveBuscada) == 0 {
		return &nodo.der
	}

	if cmp(claveBuscada, nodo.clave) < 0 {
		if nodo.izq == nil {
			return &nodo.izq
		} else {
			return nodo.izq.obtenerRefARefAHijo(claveBuscada, cmp)
		}
	} else { // esto es igual a if cmp(claveBuscada, nodo.clave) > 0, porque estas dos claves nunca pueden ser iguales dentro de esta funcion
		if nodo.der == nil {
			return &nodo.der
		} else {
			return nodo.der.obtenerRefARefAHijo(claveBuscada, cmp)
		}
	}
}

func (abb *abb[K, V]) validarRefAClave(refANodo *nodoABB[K, V]) {
	if refANodo == nil {
		panic("La clave no pertenece al diccionario")
	}
}

func (nodo *nodoABB[K, V]) estaEnRango(desde *K, hasta *K, cmp func(K, K) int) bool {
	return !(desde != nil && cmp(nodo.clave, *desde) < 0 || hasta != nil && cmp(nodo.clave, *hasta) > 0) // es !estaFueraDeRango
}

type iterABB[K comparable, V any] struct {
	cmp             func(K, K) int
	desde           *K
	hasta           *K
	pilaIteraciones TDAPila.Pila[*nodoABB[K, V]]
}

func (it *iterABB[K, V]) HaySiguiente() bool {
	return !it.pilaIteraciones.EstaVacia()
}

func (it *iterABB[K, V]) VerActual() (K, V) {
	it.validarIterador()

	nodoActual := it.pilaIteraciones.VerTope()

	return nodoActual.clave, nodoActual.dato
}

func (it *iterABB[K, V]) Siguiente() {
	it.validarIterador()

	nodoAnterior := it.pilaIteraciones.Desapilar()

	if it.hasta == nil || it.cmp(nodoAnterior.clave, *(it.hasta)) < 0 { // verifica si hay lugar entre la clave (estaba si o si en rango) y hasta ...[desde....clave......hasta]...
		it.apilarIteraciones(nodoAnterior.der)
	}
}

func (it *iterABB[K, V]) apilarIteraciones(nodo *nodoABB[K, V]) {
	if nodo == nil {
		return
	}

	if nodo.estaEnRango(it.desde, it.hasta, it.cmp) {
		it.pilaIteraciones.Apilar(nodo)
	}
	if it.desde == nil || it.cmp(*(it.desde), nodo.clave) < 0 { // ...[desde.....hasta].......clave...		o	...[desde.........clave......hasta]...
		it.apilarIteraciones(nodo.izq)
	}
	if it.desde != nil && it.cmp(nodo.clave, *(it.desde)) < 0 { // ...clave.......[desde.....hasta]...
		it.apilarIteraciones(nodo.der)
	}
}

func (it *iterABB[K, V]) validarIterador() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
