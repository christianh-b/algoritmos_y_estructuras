package lista

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento a la lista, al principio de la misma.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a la lista, al final de la misma.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma
	// y se devuelve ese valor. Si esta vacia, entra en panico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si la lista tiene elementos, se devuelve el valor del primero.
	// Si esta vacia, entra en panico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo de la lista. Si la lista tiene elementos, se devuelve el valor del ultimo.
	// Si esta vacia, entra en panico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Iterar recorre todos los elementos de la lista y en cada iteracion ejecuta la funcion "visitar" con el elemento actual como
	// parametro de entrada. Si la funcion visitar devuelve true, se avanza a la siguiente iteracion. Si devuelve false, se detiene
	// la ejecucion de Iterar.
	Iterar(visitar func(T) bool)

	// Iterador devuelve una instancia de un iterador para la lista enlazada, ubicado en la primera posicion de la misma.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el valor del elemento actual. Si la lista tiene un actual devuelve su valor.
	// Si no tiene actual, entra en panico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve true si la lista tiene mas elementos para ver, false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza el iterador al siguiente elemento de la lista.
	// Si se llama a Siguiente y el iterador ya itero todos los elementos, entra en panico con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un elemento en la posicion actual del iterador.
	Insertar(T)

	// Borrar elimina el elemento en la posicion actual del iterador.
	// Si se llama a Borrar y el iterador ya itero todos los elementos, entra en panico con un mensaje "El iterador termino de iterar.
	Borrar() T
}
