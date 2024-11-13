package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"

	"strings"

	"math/rand"

	"github.com/stretchr/testify/require"
)

const NRO_PRUEBAS = 10000

func compararInts(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	} else {
		return 0
	}
}

func TestDiccOrdenadoVacio(t *testing.T) { // (general dicc)
	t.Log("Comprueba que un Diccionario nuevo no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccOrdenadoClaveDefault(t *testing.T) { // (general dicc)
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](compararInts)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) { // (general dicc)
	t.Log("Comprueba que un Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccOrdenadoGuardar(t *testing.T) { // (general dicc)
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[1]))
	dic.Guardar(claves[1], valores[1])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[0], valores[0])
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatos(t *testing.T) { // (general dicc) (reemplazo de datos)
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func randomizarSlice[T any](arr []T) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func TestDiccOrdenadoBorrar(t *testing.T) { // (general dicc) (borrar con 1 y 2 hijos)
	t.Log("Guarda elementos en el diccionario, y se borran algunos, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Chancho"
	clave2 := "Gato"
	clave3 := "Pato"
	clave4 := "Perro"
	clave5 := "Vaca"
	valor1 := "oink"
	valor2 := "miau"
	valor3 := "quack"
	valor4 := "guau"
	valor5 := "moo"
	claves := []string{clave1, clave2, clave3, clave4, clave5}
	valores := []string{valor1, valor2, valor3, valor4, valor5}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[3], valores[3]) // "Perro": "guau"
	dic.Guardar(claves[4], valores[4]) // "Vaca": "moo"
	dic.Guardar(claves[1], valores[1]) // "Gato": "miau"
	dic.Guardar(claves[0], valores[0]) // "Chancho": "oink"
	dic.Guardar(claves[2], valores[2]) // "Pato": "quack"

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1])) // borrar con 2 hijos
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 4, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2])) // borrar con 1 hijo
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 3, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[2]) })

	require.True(t, dic.Pertenece(claves[3]))
	require.EqualValues(t, valores[3], dic.Borrar(claves[3])) // borrar raiz
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[3]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[3]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[3]) })

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.True(t, dic.Pertenece(claves[4]))
	require.EqualValues(t, valores[4], dic.Obtener(claves[4]))
}

func TestReutilizacionBorrado(t *testing.T) { // (general dicc) (re-guardar clave borrada)
	t.Log("Prueba de caja blanca: revisa que no haya problema reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestIntsComoClaves(t *testing.T) { // (general dicc) (int como clave)
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](compararInts)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestStructsComoClaves(t *testing.T) { // (general dicc) (structs como claves)
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionario.CrearABB[avanzado, int](func(p avanzado, q avanzado) int {
		if p.y.b > q.y.b {
			return 1
		}
		if p.y.b < q.y.b {
			return -1
		} else {
			return 0
		}
	})

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
}

func TestConClaveVacia(t *testing.T) { // (general dicc) (clave vacia)
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestConValorNulo(t *testing.T) { // (general dicc) (valor nulo)
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestGuardarYBorrarVolumen(t *testing.T) { // volumen
	t.Log("Esta prueba guarda y borra muchas veces, verificando el correcto funcionamiento en cada caso")
	dic := TDADiccionario.CrearABB[int, int](compararInts)

	elementos := make([]int, NRO_PRUEBAS)
	for i := 0; i < NRO_PRUEBAS; i++ {
		elementos[i] = i
	}
	randomizarSlice(elementos)

	for _, elem := range elementos {
		dic.Guardar(elem, elem)
	}
	for i := 0; i < NRO_PRUEBAS; i++ {
		require.True(t, dic.Pertenece(i))
		require.EqualValues(t, i, dic.Borrar(i))
		require.False(t, dic.Pertenece(i))
	}
}

func TestVolumen(t *testing.T) { // volumen
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := make([]string, NRO_PRUEBAS)
	valores := make([]int, NRO_PRUEBAS)

	for i := 0; i < NRO_PRUEBAS; i++ {
		valores[i] = i
	}
	randomizarSlice(valores)

	/* Inserta 'NRO_PRUEBAS' parejas en el ABB */
	for i, v := range valores {
		claves[i] = fmt.Sprintf("%d", v)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(t, NRO_PRUEBAS, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < NRO_PRUEBAS; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(t, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(t, NRO_PRUEBAS, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < NRO_PRUEBAS; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(t, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestIterarClaves(t *testing.T) { // Iterar [sin corte] (solo usa claves)
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[1], nil) // "Perro": nil
	dic.Guardar(claves[0], nil) // "Gato": nil
	dic.Guardar(claves[2], nil) // "Vaca": nil

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, _ *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.EqualValues(t, claves[0], cs[0])
	require.EqualValues(t, claves[1], cs[1])
	require.EqualValues(t, claves[2], cs[2])
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIterarValores(t *testing.T) { // Iterar [sin corte] (solo usa valores)
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Burrito"
	clave2 := "Gato"
	clave3 := "Hamster"
	clave4 := "Perro"
	clave5 := "Vaca"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave3, 3) // "Hamster": 3
	dic.Guardar(clave2, 2) // "Gato": 2
	dic.Guardar(clave1, 6) // "Burrito": 6
	dic.Guardar(clave4, 4) // "Perro": 4
	dic.Guardar(clave5, 5) // "Vaca": 5

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIterarValoresConBorrados(t *testing.T) { // Iterar [sin corte] (solo usa valores; se borraron entradas)
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Burrito"
	clave1 := "Elefante"
	clave2 := "Gato"
	clave3 := "Hamster"
	clave4 := "Perro"
	clave5 := "Vaca"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIterarConCorteVolumen(t *testing.T) { // Iterar [con corte]-volumen
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](compararInts)

	/* Inserta 'NRO_PRUEBAS' parejas en el ABB */
	elementos := make([]int, NRO_PRUEBAS)
	for i := 0; i < NRO_PRUEBAS; i++ {
		elementos[i] = i
	}
	randomizarSlice(elementos)

	for _, elem := range elementos {
		dic.Guardar(elem, elem)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c == NRO_PRUEBAS/2 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIteradorDiccOrdenado(t *testing.T) { // Iterador
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[2], valores[2])
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])

	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, claves[0], primero)

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.EqualValues(t, claves[1], segundo)
	require.EqualValues(t, valores[1], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, claves[2], tercero)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorDiccOrdenadoVacio(t *testing.T) { // Iterador (diccionario vacio)
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorVariosIteradores(t *testing.T) { // Iterador (varios iteradores en simultaneo)
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[2], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[0], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.EqualValues(t, claves[0], primero)
	require.EqualValues(t, claves[1], segundo)
	require.EqualValues(t, claves[2], tercero)
}

func TestIteradorConBorrados(t *testing.T) { // Iterador (se borraron entradas)
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del ABB y su iterador externo " +
		"tras haber borrado entradas")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave1, "")
	dic.Guardar(clave3, "")
	dic.Guardar(clave2, "")
	dic.Borrar(clave3)
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorVolumen(t *testing.T) { // Iterador-volumen
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	claves := make([]string, NRO_PRUEBAS)
	valores := make([]int, NRO_PRUEBAS)

	for i := 0; i < NRO_PRUEBAS; i++ {
		valores[i] = i
	}
	randomizarSlice(valores)

	/* Inserta 'NRO_PRUEBAS' parejas en el ABB */
	for i, v := range valores {
		claves[i] = fmt.Sprintf("%d", v)
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(t, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < NRO_PRUEBAS; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = -1
		iter.Siguiente()
	}
	require.True(t, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(t, NRO_PRUEBAS, i, "No se recorrió todo el largo")
	require.False(t, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < NRO_PRUEBAS; i++ {
		if valores[i] != -1 {
			ok = false
			break
		}
	}
	require.True(t, ok, "No se cambiaron todos los elementos")
}

func TestIterarRango(t *testing.T) { // IterarRango [sin corte] (limites de rango no pertenecen al diccionario)
	dic := TDADiccionario.CrearABB[int, int](compararInts)
	elementos := []int{100, 50, 20, 30, 120, 105, 110, 150, 125, 160}

	for _, elem := range elementos {
		dic.Guardar(elem, elem)
	}

	min := 25
	max := 130
	suma := 0
	iterados := []int{}
	dic.IterarRango(&min, &max, func(clave, dato int) bool {
		suma += clave
		iterados = append(iterados, clave)
		return true
	})

	require.EqualValues(t, 640, suma)
	require.Equal(t, []int{30, 50, 100, 105, 110, 120, 125}, iterados)
}

func TestIterarRangoDesdeMinimo(t *testing.T) { // IterarRango [sin corte] (limite "hasta" pertenece al diccionario)
	dic := TDADiccionario.CrearABB[int, int](compararInts)
	elementos := []int{100, 50, 20, 30, 120, 105, 110, 150, 125, 160}

	for _, elem := range elementos {
		dic.Guardar(elem, elem)
	}

	min := (*int)(nil)
	max := 150
	suma := 0
	iterados := []int{}
	dic.IterarRango(min, &max, func(clave, dato int) bool {
		suma += clave
		iterados = append(iterados, clave)
		return true
	})

	require.EqualValues(t, 810, suma)
	require.Equal(t, []int{20, 30, 50, 100, 105, 110, 120, 125, 150}, iterados)
}

func TestIterarRangoHastaMaximo(t *testing.T) { // IterarRango [sin corte] (limite "desde" pertenece al diccionario)
	dic := TDADiccionario.CrearABB[int, int](compararInts)
	elementos := []int{100, 50, 20, 30, 120, 105, 110, 150, 125, 160}

	for _, elem := range elementos {
		dic.Guardar(elem, elem)
	}

	min := 110
	max := (*int)(nil)
	suma := 0
	iterados := []int{}
	dic.IterarRango(&min, max, func(clave, dato int) bool {
		suma += clave
		iterados = append(iterados, clave)
		return true
	})

	require.EqualValues(t, 665, suma)
	require.Equal(t, []int{110, 120, 125, 150, 160}, iterados)
}

func TestIterarRangoConCorte(t *testing.T) { // IterarRango [con corte] (limites de rango pertenecen al diccionario)
	dic := TDADiccionario.CrearABB[int, int](compararInts)
	elementos := []int{100, 50, 20, 30, 120, 105, 110, 150, 125, 160}

	for _, elem := range elementos {
		dic.Guardar(elem, elem)
	}

	min := 30
	max := 150
	suma := 0
	iterados := []int{}
	dic.IterarRango(&min, &max, func(clave, dato int) bool {
		iterados = append(iterados, clave)
		if clave > 120 {
			return false
		}
		suma += clave
		return true
	})

	require.EqualValues(t, 515, suma)
	require.Equal(t, []int{30, 50, 100, 105, 110, 120, 125}, iterados)
}

func TestIteradorRangoVolumen(t *testing.T) { // IteradorRango-volumen
	dic := TDADiccionario.CrearABB[int, *int](compararInts)

	valores := make([]int, NRO_PRUEBAS)

	for i := 0; i < NRO_PRUEBAS; i++ {
		valores[i] = i
	}
	randomizarSlice(valores)

	/* Inserta 'NRO_PRUEBAS' parejas en el ABB */
	for i, v := range valores {
		dic.Guardar(v, &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	min := 100
	max := 1000
	iter := dic.IteradorRango(&min, &max)
	require.True(t, iter.HaySiguiente())

	iterados := []int{}
	iteradosEsperados := []int{}

	for iter.HaySiguiente() {
		claveActual, _ := iter.VerActual()
		iterados = append(iterados, claveActual)
		iter.Siguiente()
	}

	for i := min; i <= max; i++ {
		iteradosEsperados = append(iteradosEsperados, i)
	}

	require.False(t, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")
	require.Equal(t, iteradosEsperados, iterados)
}
