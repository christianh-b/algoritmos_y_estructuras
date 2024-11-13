package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola debe estar vacia al momento de crearla")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, " '''Panic''' No se puede ver el primer elemento de una cola que esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, " '''Panic''' No se puede desencolar elementos de una cola que esta vacia")

}

func TestInvarianteFIFO(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	require.Equal(t, 1, cola.Desencolar(), "No mantiene el invariante FIFO; el primer elemento desencolado debe ser igual al primer elemento que se encolo")
	require.Equal(t, 2, cola.Desencolar(), "No mantiene el invariante FIFO; el segundo elemento desencolado debe ser igual al segundo elemento que se encolo")
	require.Equal(t, 3, cola.Desencolar(), "No mantiene el invariante FIFO; el tercer elemento desencolado debe ser igual al tercer elemento que se encolo")

}

func TestApilarDiferentesTipos(t *testing.T) {
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	colaStrigns := TDACola.CrearColaEnlazada[string]()
	colaBools := TDACola.CrearColaEnlazada[bool]()

	colaEnteros.Encolar(1)
	colaEnteros.Encolar(2)
	require.Equal(t, 1, colaEnteros.Desencolar(), "No mantiene el invariante FIFO; el primer elemento desencolado debe ser igual al primer elemento que se encolo")
	require.Equal(t, 2, colaEnteros.Desencolar(), "No mantiene el invariante FIFO; el segundo elemento desencolado debe ser igual al segundo elemento que se encolo")

	colaStrigns.Encolar("Hola")
	colaStrigns.Encolar("Mundo")
	require.Equal(t, "Hola", colaStrigns.Desencolar(), "No mantiene el invariante FIFO; el primer elemento desencolado debe ser igual al primer elemento que se encolo")
	require.Equal(t, "Mundo", colaStrigns.Desencolar(), "No mantiene el invariante FIFO; el segundo elemento desencolado debe ser igual al segundo elemento que se encolo")

	colaBools.Encolar(true)
	colaBools.Encolar(false)
	require.Equal(t, true, colaBools.Desencolar(), "No mantiene el invariante FIFO; el primer elemento desencolado debe ser igual al primer elemento que se encolo")
	require.Equal(t, false, colaBools.Desencolar(), "No mantiene el invariante FIFO; el segundo elemento desencolado debe ser igual al segundo elemento que se encolo")

	require.True(t, colaEnteros.EstaVacia(), "La cola debe estar vacia luego de desencolar todos los elementos que contenia")
	require.True(t, colaStrigns.EstaVacia(), "La cola debe estar vacia luego de desencolar todos los elementos que contenia")
	require.True(t, colaBools.EstaVacia(), "La cola debe estar vacia luego de desencolar todos los elementos que contenia")

	require.PanicsWithValue(t, "La cola esta vacia", func() { colaEnteros.VerPrimero() }, " '''Panic''' No se puede ver el primer elemento de una cola que esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaEnteros.Desencolar() }, " '''Panic''' No se puede desencolar elementos de una cola que esta vacia")

	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStrigns.VerPrimero() }, " '''Panic''' No se puede ver el primer elemento de una cola que esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaStrigns.Desencolar() }, " '''Panic''' No se puede desencolar elementos de una cola que esta vacia")

	require.PanicsWithValue(t, "La cola esta vacia", func() { colaBools.VerPrimero() }, " '''Panic''' No se puede ver el primer elemento de una cola que esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaBools.Desencolar() }, " '''Panic''' No se puede desencolar elementos de una cola que esta vacia")
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	var elementos = 1000
	for i := 1; i <= elementos; i++ {
		cola.Encolar(i)
	}
	for i := 1; i <= elementos; i++ {
		elementoEsperado := i
		elementoObtenido := cola.Desencolar()
		require.Equal(t, elementoEsperado, elementoObtenido, "El elemento desencolado no es el esperado despues de desencolar %d elementos", elementos-i+1)
	}
	require.True(t, cola.EstaVacia(), "La cola no esta vacia despues de desencolar todos los elementos")

	elementos = 10000
	for i := 1; i <= elementos; i++ {
		cola.Encolar(i)
	}
	for i := 1; i <= elementos; i++ {
		elementoEsperado := i
		elementoObtenido := cola.Desencolar()
		require.Equal(t, elementoEsperado, elementoObtenido, "El elemento desencolado no es el esperado despues de desencolar %d elementos", elementos-i+1)
	}
	require.True(t, cola.EstaVacia(), "La cola no esta vacia despues de desencolar todos los elementos")
}
