package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "La pila debe estar vacia al momento de crearla")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, " '''Panic''' No se puede ver el tope de una pila que esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, " '''Panic''' No se puede desapilar elementos de una pila que esta vacia")

}

func TestInvarianteLIFO(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	require.Equal(t, 2, pila.Desapilar(), "No mantiene el invariante LIFO; el elemento desapilado debe ser igual al ultimo elemento que se apilo")
	require.Equal(t, 1, pila.Desapilar(), "No mantiene el invariante LIFO; el elemento desapilado debe ser igual al ultimo elemento que se apilo")

}

func TestApilarDiferentesTipos(t *testing.T) {
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaStrings := TDAPila.CrearPilaDinamica[string]()
	pilaBools := TDAPila.CrearPilaDinamica[bool]()

	pilaEnteros.Apilar(1)
	pilaEnteros.Apilar(2)
	require.Equal(t, 2, pilaEnteros.Desapilar(), "No mantiene el invariante LIFO; el elemento desapilado debe ser igual al ultimo elemento que se apilo")
	require.Equal(t, 1, pilaEnteros.Desapilar(), "No mantiene el invariante LIFO; el elemento desapilado debe ser igual al ultimo elemento que se apilo")

	pilaStrings.Apilar("Hola")
	pilaStrings.Apilar("Mundo")
	require.Equal(t, "Mundo", pilaStrings.Desapilar(), "No mantiene el invariante LIFO; el elemento desapilado debe ser igual al ultimo elemento que se apilo")
	require.Equal(t, "Hola", pilaStrings.Desapilar(), "No mantiene el invariante LIFO; el elemento desapilado debe ser igual al ultimo elemento que se apilo")

	pilaBools.Apilar(true)
	pilaBools.Apilar(false)
	require.Equal(t, false, pilaBools.Desapilar(), "No mantiene el invariante LIFO; el elemento desapilado debe ser igual al ultimo elemento que se apilo")
	require.Equal(t, true, pilaBools.Desapilar(), "No mantiene el invariante LIFO; el elemento desapilado debe ser igual al ultimo elemento que se apilo")

	require.True(t, pilaEnteros.EstaVacia(), "La pila debe estar vacia luego de desapilar todos los elementos que contenia")
	require.True(t, pilaStrings.EstaVacia(), "La pila debe estar vacia luego de desapilar todos los elementos que contenia")
	require.True(t, pilaBools.EstaVacia(), "La pila debe estar vacia luego de desapilar todos los elementos que contenia")

	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaEnteros.VerTope() }, " '''Panic''' No se puede ver el tope de una pila que esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaEnteros.Desapilar() }, " '''Panic''' No se puede desapilar elementos de una pila que esta vacia")

	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStrings.VerTope() }, " '''Panic''' No se puede ver el tope de una pila que esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaStrings.Desapilar() }, " '''Panic''' No se puede desapilar elementos de una pila que esta vacia")

	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaBools.VerTope() }, " '''Panic''' No se puede ver el tope de una pila que esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaBools.Desapilar() }, " '''Panic''' No se puede desapilar elementos de una pila que esta vacia")
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	var elementos = 1000
	for i := 1; i <= elementos; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope(), "El tope de la pila no es el esperado despues de apilar %d elementos", i)
	}
	for i := elementos; i > 0; i-- {
		require.Equal(t, i, pila.VerTope(), "El tope de la pila no es el esperado antes de desapilar %d elementos", elementos-i+1)
		require.Equal(t, i, pila.Desapilar(), "El elemento desapilado no es el esperado despues de desapilar %d elementos", elementos-i+1)
	}
	require.True(t, pila.EstaVacia(), "La pila no esta vacia despues de desapilar todos los elementos")

	elementos = 10000
	for i := 1; i <= elementos; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope(), "El tope de la pila no es el esperado despues de apilar %d elementos", i)
	}
	for i := elementos; i > 0; i-- {
		require.Equal(t, i, pila.VerTope(), "El tope de la pila no es el esperado antes de desapilar %d elementos", elementos-i+1)
		require.Equal(t, i, pila.Desapilar(), "El elemento desapilado no es el esperado despues de desapilar %d elementos", elementos-i+1)
	}
	require.True(t, pila.EstaVacia(), "La pila no esta vacia despues de desapilar todos los elementos")
}
