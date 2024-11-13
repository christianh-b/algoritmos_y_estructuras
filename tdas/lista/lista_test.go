package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const volPruebas int = 10000

func TestListaNueva(t *testing.T) { // 1, 5 y 6
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaString := TDALista.CrearListaEnlazada[string]()
	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaFloat64 := TDALista.CrearListaEnlazada[float64]()

	require.True(t, listaInt.EstaVacia())
	require.Equal(t, 0, listaInt.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaInt.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaInt.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaInt.BorrarPrimero() })

	require.True(t, listaString.EstaVacia())
	require.Equal(t, 0, listaString.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaString.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaString.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaString.BorrarPrimero() })

	require.True(t, listaBool.EstaVacia())
	require.Equal(t, 0, listaBool.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaBool.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaBool.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaBool.BorrarPrimero() })

	require.True(t, listaFloat64.EstaVacia())
	require.Equal(t, 0, listaFloat64.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFloat64.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFloat64.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFloat64.BorrarPrimero() })
}

func TestInsertarPrimero(t *testing.T) { // 2
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaString := TDALista.CrearListaEnlazada[string]()
	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaFloat64 := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarPrimero(1)
	require.False(t, listaInt.EstaVacia())
	require.Equal(t, 1, listaInt.Largo())
	require.Equal(t, 1, listaInt.VerPrimero())
	require.Equal(t, 1, listaInt.VerUltimo())
	listaInt.InsertarPrimero(0)
	require.Equal(t, 0, listaInt.VerPrimero())
	require.Equal(t, 1, listaInt.VerUltimo())

	listaString.InsertarPrimero("uno")
	require.False(t, listaString.EstaVacia())
	require.Equal(t, 1, listaString.Largo())
	require.Equal(t, "uno", listaString.VerPrimero())
	require.Equal(t, "uno", listaString.VerUltimo())
	listaString.InsertarPrimero("cero")
	require.Equal(t, "cero", listaString.VerPrimero())
	require.Equal(t, "uno", listaString.VerUltimo())

	listaBool.InsertarPrimero(true)
	require.False(t, listaBool.EstaVacia())
	require.Equal(t, 1, listaBool.Largo())
	require.Equal(t, true, listaBool.VerPrimero())
	require.Equal(t, true, listaBool.VerUltimo())
	listaBool.InsertarPrimero(false)
	require.Equal(t, false, listaBool.VerPrimero())
	require.Equal(t, true, listaBool.VerUltimo())

	listaFloat64.InsertarPrimero(3.14)
	require.False(t, listaFloat64.EstaVacia())
	require.Equal(t, 1, listaFloat64.Largo())
	require.Equal(t, 3.14, listaFloat64.VerPrimero())
	require.Equal(t, 3.14, listaFloat64.VerUltimo())
	listaFloat64.InsertarPrimero(2.71828)
	require.Equal(t, 2.71828, listaFloat64.VerPrimero())
	require.Equal(t, 3.14, listaFloat64.VerUltimo())
}

func TestInsertarUltimo(t *testing.T) { // 2
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaString := TDALista.CrearListaEnlazada[string]()
	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaFloat64 := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarUltimo(1)
	require.False(t, listaInt.EstaVacia())
	require.Equal(t, 1, listaInt.Largo())
	require.Equal(t, 1, listaInt.VerPrimero())
	require.Equal(t, 1, listaInt.VerUltimo())
	listaInt.InsertarUltimo(2)
	require.Equal(t, 1, listaInt.VerPrimero())
	require.Equal(t, 2, listaInt.VerUltimo())

	listaString.InsertarUltimo("uno")
	require.False(t, listaString.EstaVacia())
	require.Equal(t, 1, listaString.Largo())
	require.Equal(t, "uno", listaString.VerPrimero())
	require.Equal(t, "uno", listaString.VerUltimo())
	listaString.InsertarUltimo("dos")
	require.Equal(t, "uno", listaString.VerPrimero())
	require.Equal(t, "dos", listaString.VerUltimo())

	listaBool.InsertarUltimo(true)
	require.False(t, listaBool.EstaVacia())
	require.Equal(t, 1, listaBool.Largo())
	require.Equal(t, true, listaBool.VerPrimero())
	require.Equal(t, true, listaBool.VerUltimo())
	listaBool.InsertarUltimo(false)
	require.Equal(t, true, listaBool.VerPrimero())
	require.Equal(t, false, listaBool.VerUltimo())

	listaFloat64.InsertarUltimo(3.14)
	require.False(t, listaFloat64.EstaVacia())
	require.Equal(t, 1, listaFloat64.Largo())
	require.Equal(t, 3.14, listaFloat64.VerPrimero())
	require.Equal(t, 3.14, listaFloat64.VerUltimo())
	listaFloat64.InsertarUltimo(2.71828)
	require.Equal(t, 3.14, listaFloat64.VerPrimero())
	require.Equal(t, 2.71828, listaFloat64.VerUltimo())
}

func TestBorrarPrimero(t *testing.T) { // 2
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaString := TDALista.CrearListaEnlazada[string]()
	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaFloat64 := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarUltimo(-1)
	listaInt.InsertarUltimo(-2)
	require.Equal(t, 2, listaInt.Largo())
	require.Equal(t, -1, listaInt.BorrarPrimero())
	require.Equal(t, -2, listaInt.VerPrimero())
	require.Equal(t, 1, listaInt.Largo())

	listaString.InsertarUltimo("hola")
	listaString.InsertarUltimo("chau")
	require.Equal(t, 2, listaString.Largo())
	require.Equal(t, "hola", listaString.BorrarPrimero())
	require.Equal(t, "chau", listaString.VerPrimero())
	require.Equal(t, 1, listaString.Largo())

	listaBool.InsertarUltimo(true)
	listaBool.InsertarUltimo(false)
	require.Equal(t, 2, listaBool.Largo())
	require.Equal(t, true, listaBool.BorrarPrimero())
	require.Equal(t, false, listaBool.VerPrimero())
	require.Equal(t, 1, listaBool.Largo())

	listaFloat64.InsertarUltimo(0.32)
	listaFloat64.InsertarUltimo(0.64)
	require.Equal(t, 2, listaFloat64.Largo())
	require.Equal(t, 0.32, listaFloat64.BorrarPrimero())
	require.Equal(t, 0.64, listaFloat64.VerPrimero())
	require.Equal(t, 1, listaFloat64.Largo())
}

func TestBorrarHastaVacia(t *testing.T) { // 4 y 7
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaString := TDALista.CrearListaEnlazada[string]()
	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaFloat64 := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarPrimero(777)
	listaInt.BorrarPrimero()
	require.True(t, listaInt.EstaVacia())
	require.Equal(t, 0, listaInt.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaInt.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaInt.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaInt.BorrarPrimero() })

	listaString.InsertarPrimero("setenta")
	listaString.BorrarPrimero()
	require.True(t, listaString.EstaVacia())
	require.Equal(t, 0, listaString.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaString.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaString.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaString.BorrarPrimero() })

	listaBool.InsertarPrimero(false)
	listaBool.BorrarPrimero()
	require.True(t, listaBool.EstaVacia())
	require.Equal(t, 0, listaBool.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaBool.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaBool.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaBool.BorrarPrimero() })

	listaFloat64.InsertarPrimero(-0.123456789)
	listaFloat64.BorrarPrimero()
	require.True(t, listaFloat64.EstaVacia())
	require.Equal(t, 0, listaFloat64.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFloat64.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFloat64.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFloat64.BorrarPrimero() })
}

func TestVolumen(t *testing.T) { // 3
	n := volPruebas / 2
	lista := TDALista.CrearListaEnlazada[int]()
	cont := 0

	for i := n; i < volPruebas; i++ {
		lista.InsertarUltimo(i)
		cont++
		require.Equal(t, n, lista.VerPrimero())
		require.Equal(t, i, lista.VerUltimo())
		require.Equal(t, cont, lista.Largo())
	}

	for i := n - 1; i >= 0; i-- {
		lista.InsertarPrimero(i)
		cont++
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, volPruebas-1, lista.VerUltimo())
		require.Equal(t, cont, lista.Largo())
	}

	for i := range volPruebas - 1 {
		require.Equal(t, i, lista.BorrarPrimero())
		cont--
		require.Equal(t, i+1, lista.VerPrimero())
		require.Equal(t, volPruebas-1, lista.VerUltimo())
		require.Equal(t, cont, lista.Largo())
	}

	require.Equal(t, volPruebas-1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.Iterar(contar)
	require.Equal(t, 0, cant)

	for i := range volPruebas + 1 {
		lista.InsertarUltimo(i)
	}

	lista.Iterar(contar)
	require.Equal(t, volPruebas+1, cant)

	lista.Iterar(sumar)
	require.Equal(t, volPruebas*(volPruebas+1)/2, suma)

	lista.Iterar(sumar5)
	require.Equal(t, 10, suma5)
}

var cant int

func contar(v int) bool {
	cant += 1
	return true
}

var suma int

func sumar(v int) bool {
	suma += v
	return true
}

var suma5 int

func sumar5(v int) bool {
	if v == 5 {
		return false
	}
	suma5 += v
	return true
}

func TestIteradorExternoRecienCreado(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaStr := TDALista.CrearListaEnlazada[string]()
	listaFloat := TDALista.CrearListaEnlazada[float64]()

	itInt := listaInt.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itInt.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itInt.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itInt.Borrar() })

	itStr := listaStr.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itStr.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itStr.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itStr.Borrar() })

	itFloat := listaFloat.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itFloat.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itFloat.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itFloat.Borrar() })
}

func TestIteradorExternoInsertarAlFinal(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaStr := TDALista.CrearListaEnlazada[string]()
	listaFloat := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarPrimero(1)
	itInt := listaInt.Iterador()
	require.Equal(t, 1, itInt.VerActual())
	itInt.Siguiente()
	require.False(t, itInt.HaySiguiente())
	itInt.Insertar(2)
	require.Equal(t, 2, itInt.VerActual())
	itInt.Siguiente()
	require.False(t, itInt.HaySiguiente())
	itInt.Insertar(3)
	require.Equal(t, 3, itInt.VerActual())
	itInt.Siguiente()
	require.False(t, itInt.HaySiguiente())

	listaStr.InsertarPrimero("uno")
	itStr := listaStr.Iterador()
	require.Equal(t, "uno", itStr.VerActual())
	itStr.Siguiente()
	require.False(t, itStr.HaySiguiente())
	itStr.Insertar("dos")
	require.Equal(t, "dos", itStr.VerActual())
	itStr.Siguiente()
	require.False(t, itStr.HaySiguiente())
	itStr.Insertar("tres")
	require.Equal(t, "tres", itStr.VerActual())
	itStr.Siguiente()
	require.False(t, itStr.HaySiguiente())

	listaFloat.InsertarPrimero(1.1)
	itFloat := listaFloat.Iterador()
	require.Equal(t, 1.1, itFloat.VerActual())
	itFloat.Siguiente()
	require.False(t, itFloat.HaySiguiente())
	itFloat.Insertar(2.2)
	require.Equal(t, 2.2, itFloat.VerActual())
	itFloat.Siguiente()
	require.False(t, itFloat.HaySiguiente())
	itFloat.Insertar(3.3)
	require.Equal(t, 3.3, itFloat.VerActual())
	itFloat.Siguiente()
	require.False(t, itFloat.HaySiguiente())
}

func TestIteradorExternoInsertarEnMedio(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaStr := TDALista.CrearListaEnlazada[string]()
	listaFloat := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarPrimero(1)
	listaInt.InsertarUltimo(3)
	itInt := listaInt.Iterador()
	require.Equal(t, 1, itInt.VerActual())
	require.True(t, itInt.HaySiguiente())
	itInt.Siguiente()
	itInt.Insertar(2)
	require.Equal(t, 2, itInt.VerActual())
	require.True(t, itInt.HaySiguiente())
	itInt.Siguiente()
	require.Equal(t, 3, itInt.VerActual())
	itInt.Siguiente()
	require.False(t, itInt.HaySiguiente())

	listaStr.InsertarPrimero("uno")
	listaStr.InsertarUltimo("tres")
	itStr := listaStr.Iterador()
	require.Equal(t, "uno", itStr.VerActual())
	require.True(t, itStr.HaySiguiente())
	itStr.Siguiente()
	itStr.Insertar("dos")
	require.Equal(t, "dos", itStr.VerActual())
	require.True(t, itStr.HaySiguiente())
	itStr.Siguiente()
	require.Equal(t, "tres", itStr.VerActual())
	itStr.Siguiente()
	require.False(t, itStr.HaySiguiente())

	listaFloat.InsertarPrimero(1.1)
	listaFloat.InsertarUltimo(3.3)
	itFloat := listaFloat.Iterador()
	require.Equal(t, 1.1, itFloat.VerActual())
	require.True(t, itFloat.HaySiguiente())
	itFloat.Siguiente()
	itFloat.Insertar(2.2)
	require.Equal(t, 2.2, itFloat.VerActual())
	require.True(t, itFloat.HaySiguiente())
	itFloat.Siguiente()
	require.Equal(t, 3.3, itFloat.VerActual())
	itFloat.Siguiente()
	require.False(t, itFloat.HaySiguiente())
}

func TestIteradorExternoRemoverPrimero(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaStr := TDALista.CrearListaEnlazada[string]()
	listaFloat := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarPrimero(1)
	listaInt.InsertarUltimo(2)
	listaInt.InsertarUltimo(3)
	require.Equal(t, 3, listaInt.Largo())
	itInt := listaInt.Iterador()
	require.Equal(t, 1, listaInt.VerPrimero())
	itInt.Borrar()
	require.Equal(t, 2, listaInt.VerPrimero())
	require.Equal(t, 2, listaInt.Largo())
	itInt.Borrar()
	require.Equal(t, 3, listaInt.VerPrimero())
	require.Equal(t, 1, listaInt.Largo())

	listaStr.InsertarPrimero("uno")
	listaStr.InsertarUltimo("dos")
	listaStr.InsertarUltimo("tres")
	require.Equal(t, 3, listaStr.Largo())
	itStr := listaStr.Iterador()
	require.Equal(t, "uno", listaStr.VerPrimero())
	itStr.Borrar()
	require.Equal(t, "dos", listaStr.VerPrimero())
	require.Equal(t, 2, listaStr.Largo())
	itStr.Borrar()
	require.Equal(t, "tres", listaStr.VerPrimero())
	require.Equal(t, 1, listaStr.Largo())

	listaFloat.InsertarPrimero(1.1)
	listaFloat.InsertarUltimo(2.2)
	listaFloat.InsertarUltimo(3.3)
	require.Equal(t, 3, listaFloat.Largo())
	itFloat := listaFloat.Iterador()
	require.Equal(t, 1.1, listaFloat.VerPrimero())
	itFloat.Borrar()
	require.Equal(t, 2.2, listaFloat.VerPrimero())
	require.Equal(t, 2, listaFloat.Largo())
	itFloat.Borrar()
	require.Equal(t, 3.3, listaFloat.VerPrimero())
	require.Equal(t, 1, listaFloat.Largo())

}

func TestIteradorExternoRemoverUltimo(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaStr := TDALista.CrearListaEnlazada[string]()
	listaFloat := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarPrimero(1)
	listaInt.InsertarUltimo(2)
	listaInt.InsertarUltimo(3)
	require.Equal(t, 3, listaInt.Largo())
	itInt := listaInt.Iterador()
	itInt.Siguiente()
	itInt.Siguiente()
	itInt.Borrar()
	require.Equal(t, 1, listaInt.VerPrimero())
	require.Equal(t, 2, listaInt.VerUltimo())
	require.Equal(t, 2, listaInt.Largo())

	listaStr.InsertarPrimero("uno")
	listaStr.InsertarUltimo("dos")
	listaStr.InsertarUltimo("tres")
	require.Equal(t, 3, listaStr.Largo())
	itStr := listaStr.Iterador()
	itStr.Siguiente()
	itStr.Siguiente()
	itStr.Borrar()
	require.Equal(t, "uno", listaStr.VerPrimero())
	require.Equal(t, "dos", listaStr.VerUltimo())
	require.Equal(t, 2, listaStr.Largo())

	listaFloat.InsertarPrimero(1.1)
	listaFloat.InsertarUltimo(2.2)
	listaFloat.InsertarUltimo(3.3)
	require.Equal(t, 3, listaFloat.Largo())
	itFloat := listaFloat.Iterador()
	itFloat.Siguiente()
	itFloat.Siguiente()
	itFloat.Borrar()
	require.Equal(t, 1.1, listaFloat.VerPrimero())
	require.Equal(t, 2.2, listaFloat.VerUltimo())
	require.Equal(t, 2, listaFloat.Largo())
}

func TestIteradorExternoRemoverDelMedio(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaStr := TDALista.CrearListaEnlazada[string]()
	listaFloat := TDALista.CrearListaEnlazada[float64]()

	listaInt.InsertarPrimero(1)
	listaInt.InsertarUltimo(2)
	listaInt.InsertarUltimo(3)
	require.Equal(t, 3, listaInt.Largo())
	itInt := listaInt.Iterador()
	itInt.Siguiente()
	require.Equal(t, 2, itInt.VerActual())
	itInt.Borrar()
	require.Equal(t, 3, itInt.VerActual())
	itInt.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itInt.VerActual() })
	require.Equal(t, 1, listaInt.VerPrimero())
	require.Equal(t, 3, listaInt.VerUltimo())
	require.Equal(t, 2, listaInt.Largo())

	listaStr.InsertarPrimero("uno")
	listaStr.InsertarUltimo("dos")
	listaStr.InsertarUltimo("tres")
	require.Equal(t, 3, listaStr.Largo())
	itStr := listaStr.Iterador()
	itStr.Siguiente()
	require.Equal(t, "dos", itStr.VerActual())
	itStr.Borrar()
	require.Equal(t, "tres", itStr.VerActual())
	itStr.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itStr.VerActual() })
	require.Equal(t, "uno", listaStr.VerPrimero())
	require.Equal(t, "tres", listaStr.VerUltimo())
	require.Equal(t, 2, listaStr.Largo())

	listaFloat.InsertarPrimero(1.1)
	listaFloat.InsertarUltimo(2.2)
	listaFloat.InsertarUltimo(3.3)
	require.Equal(t, 3, listaFloat.Largo())
	itFloat := listaFloat.Iterador()
	itFloat.Siguiente()
	require.Equal(t, 2.2, itFloat.VerActual())
	itFloat.Borrar()
	require.Equal(t, 3.3, itFloat.VerActual())
	itFloat.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { itFloat.VerActual() })
	require.Equal(t, 1.1, listaFloat.VerPrimero())
	require.Equal(t, 3.3, listaFloat.VerUltimo())
	require.Equal(t, 2, listaFloat.Largo())
}
