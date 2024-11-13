package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const volPruebas = 100000

func TestPilaNueva(t *testing.T) { // 1, 5 y 6
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaFloat64 := TDAPila.CrearPilaDinamica[float64]()

	require.True(t, pilaInt.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.Desapilar() })

	require.True(t, pilaString.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaString.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaString.Desapilar() })

	require.True(t, pilaBool.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaBool.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaBool.Desapilar() })

	require.True(t, pilaFloat64.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaFloat64.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaFloat64.Desapilar() })
}

func TestApilar(t *testing.T) { // 2
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaFloat64 := TDAPila.CrearPilaDinamica[float64]()

	require.NotPanics(t, func() { pilaInt.Apilar(1) })
	require.False(t, pilaInt.EstaVacia())
	require.Equal(t, 1, pilaInt.VerTope())
	pilaInt.Apilar(2)
	require.Equal(t, 2, pilaInt.VerTope())

	require.NotPanics(t, func() { pilaString.Apilar("uno") })
	require.False(t, pilaString.EstaVacia())
	require.Equal(t, "uno", pilaString.VerTope())
	pilaString.Apilar("dos")
	require.Equal(t, "dos", pilaString.VerTope())

	require.NotPanics(t, func() { pilaBool.Apilar(true) })
	require.False(t, pilaBool.EstaVacia())
	require.Equal(t, true, pilaBool.VerTope())
	pilaBool.Apilar(false)
	require.Equal(t, false, pilaBool.VerTope())

	require.NotPanics(t, func() { pilaFloat64.Apilar(3.14) })
	require.False(t, pilaFloat64.EstaVacia())
	require.Equal(t, 3.14, pilaFloat64.VerTope())
	pilaFloat64.Apilar(2.71828)
	require.Equal(t, 2.71828, pilaFloat64.VerTope())
}

func TestDesapilar(t *testing.T) { // 2
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaFloat64 := TDAPila.CrearPilaDinamica[float64]()

	pilaInt.Apilar(1)
	pilaInt.Apilar(2)
	require.Equal(t, 2, pilaInt.Desapilar())
	require.Equal(t, 1, pilaInt.VerTope())

	pilaString.Apilar("hola")
	pilaString.Apilar("chau")
	require.Equal(t, "chau", pilaString.Desapilar())
	require.Equal(t, "hola", pilaString.VerTope())

	pilaBool.Apilar(false)
	pilaBool.Apilar(true)
	require.Equal(t, true, pilaBool.Desapilar())
	require.Equal(t, false, pilaBool.VerTope())

	pilaFloat64.Apilar(22.04)
	pilaFloat64.Apilar(1.22)
	require.Equal(t, 1.22, pilaFloat64.Desapilar())
	require.Equal(t, 22.04, pilaFloat64.VerTope())
}

func TestDesapilarHastaVacia(t *testing.T) { // 4 y 7
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaFloat64 := TDAPila.CrearPilaDinamica[float64]()

	pilaInt.Apilar(444)
	pilaInt.Desapilar()
	require.True(t, pilaInt.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.Desapilar() })

	pilaString.Apilar("cuarenta")
	pilaString.Desapilar()
	require.True(t, pilaString.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaString.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaString.Desapilar() })

	pilaBool.Apilar(true)
	pilaBool.Desapilar()
	require.True(t, pilaBool.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaBool.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaBool.Desapilar() })

	pilaFloat64.Apilar(0.0)
	pilaFloat64.Desapilar()
	require.True(t, pilaFloat64.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaFloat64.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaFloat64.Desapilar() })
}

func TestVolumen(t *testing.T) { // 3
	n := volPruebas
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := range n {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}

	for i := n - 1; i > 0; i-- {
		pila.Desapilar()
		require.Equal(t, i-1, pila.VerTope())
	}

	pila.Desapilar()
	require.True(t, pila.EstaVacia())
}
