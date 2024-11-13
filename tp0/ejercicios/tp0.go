package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x

}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {

	if len(vector) == 0 {
		return -1
	}

	maximo := vector[0]
	indiceDelMaximo := 0

	for i := 1; i < len(vector); i++ {
		if vector[i] > maximo {
			maximo = vector[i]
			indiceDelMaximo = i
		}
	}

	return indiceDelMaximo
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.

// [1 2] | [3 2]
func Comparar(vector1 []int, vector2 []int) int {
	longitudVector1 := len(vector1)
	longitudVector2 := len(vector2)
	longitudMinima := longitudVector1

	if longitudVector2 < longitudMinima {
		longitudMinima = longitudVector2
	}

	for i := 0; i < longitudMinima; i++ {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}

	if longitudVector1 < longitudVector2 {
		return -1
	} else if longitudVector2 < longitudVector1 {
		return 1
	}

	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	longitudArray := len(vector)

	for i := longitudArray - 1; i > 0; i-- {
		indiceDelMaximo := Maximo(vector[:i+1])

		Swap(&vector[i], &vector[indiceDelMaximo])
	}

}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return SumaRecursiva(vector, 0, len(vector), 0)
}

func SumaRecursiva(vector []int, i int, fin int, acum int) int {
	if i == fin {
		return acum
	}
	acum += vector[i]
	return SumaRecursiva(vector, i+1, fin, acum)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	if len(cadena) == 1 {
		return true
	}
	return EsCadenaCapicuaRecursiva(cadena, 0, len(cadena)-1)
}

func EsCadenaCapicuaRecursiva(cadena string, ini int, fin int) bool {
	if ini >= fin {
		return true
	}

	if cadena[ini] != cadena[fin] {
		return false
	}

	return EsCadenaCapicuaRecursiva(cadena, ini+1, fin-1)
}
