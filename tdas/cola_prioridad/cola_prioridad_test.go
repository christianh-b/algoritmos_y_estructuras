package cola_prioridad_test

import (
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapDesdeArregloVacio(t *testing.T) {
	heapDesdeArr := TDAHeap.CrearHeapArr[int]([]int{}, func(a, b int) int { return a - b })

	require.EqualValues(t, 0, heapDesdeArr.Cantidad())
	require.True(t, heapDesdeArr.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapDesdeArr.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapDesdeArr.Desencolar() })
}

func TestEncolarUnElemento(t *testing.T) {
	heapStr := TDAHeap.CrearHeap[string](strings.Compare)
	heapInt := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	heapStr.Encolar("A")

	require.EqualValues(t, 1, heapStr.Cantidad())
	require.EqualValues(t, "A", heapStr.VerMax())
	require.False(t, heapStr.EstaVacia())

	heapInt.Encolar(99)

	require.EqualValues(t, 1, heapInt.Cantidad())
	require.EqualValues(t, 99, heapInt.VerMax())
	require.False(t, heapInt.EstaVacia())

}

func TestHeapDesdeArregloConUnElemento(t *testing.T) {
	heapStr := TDAHeap.CrearHeapArr[string]([]string{"A"}, strings.Compare)
	heapInt := TDAHeap.CrearHeapArr[int]([]int{99}, func(a, b int) int { return a - b })

	require.EqualValues(t, 1, heapStr.Cantidad())
	require.EqualValues(t, "A", heapStr.VerMax())
	require.False(t, heapStr.EstaVacia())

	require.EqualValues(t, 1, heapInt.Cantidad())
	require.EqualValues(t, 99, heapInt.VerMax())
	require.False(t, heapInt.EstaVacia())
}

func TestEncolarVariosElementos(t *testing.T) {
	strs := []string{"R", "B", "S"}
	nums := []int{1, -1, 3}

	heapStr := TDAHeap.CrearHeap[string](strings.Compare)
	require.True(t, heapStr.EstaVacia())
	require.Equal(t, 0, heapStr.Cantidad())

	heapInt := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })
	require.True(t, heapInt.EstaVacia())
	require.Equal(t, 0, heapInt.Cantidad())

	heapStr.Encolar(strs[0])
	require.EqualValues(t, 1, heapStr.Cantidad())
	require.EqualValues(t, strs[0], heapStr.VerMax())

	heapInt.Encolar(nums[0])
	require.EqualValues(t, 1, heapInt.Cantidad())
	require.EqualValues(t, nums[0], heapInt.VerMax())

	heapStr.Encolar(strs[1])
	require.EqualValues(t, 2, heapStr.Cantidad())
	require.EqualValues(t, strs[0], heapStr.VerMax())

	heapInt.Encolar(nums[1])
	require.EqualValues(t, 2, heapInt.Cantidad())
	require.EqualValues(t, nums[0], heapInt.VerMax())

	heapStr.Encolar(strs[2])
	require.EqualValues(t, 3, heapStr.Cantidad())
	require.EqualValues(t, strs[2], heapStr.VerMax())

	heapInt.Encolar(nums[2])
	require.EqualValues(t, 3, heapInt.Cantidad())
	require.EqualValues(t, nums[2], heapInt.VerMax())

	require.False(t, heapStr.EstaVacia())
	require.False(t, heapInt.EstaVacia())
}

func TestHeapDesdeArregloConVariosElementos(t *testing.T) {
	strs := []string{"R", "B", "S", "A", "D", "F"}
	nums := []int{1, -1, 3, 43, 40, 100}

	heapStr := TDAHeap.CrearHeapArr[string](strs, strings.Compare)
	heapInt := TDAHeap.CrearHeapArr[int](nums, func(a, b int) int { return a - b })

	require.Equal(t, len(strs), heapStr.Cantidad())
	require.Equal(t, strs[2], heapStr.VerMax())
	require.False(t, heapStr.EstaVacia())

	require.Equal(t, len(nums), heapInt.Cantidad())
	require.Equal(t, nums[5], heapInt.VerMax())
	require.False(t, heapInt.EstaVacia())
}

func TestDesencolarVariosElementos(t *testing.T) {
	strs := []string{"A", "B", "C", "D", "E"}
	nums := []int{1, 2, 3, 4, 5}

	heapStr := TDAHeap.CrearHeap[string](strings.Compare)
	heapInt := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	heapStr.Encolar(strs[0])
	heapStr.Encolar(strs[1])
	heapStr.Encolar(strs[2])
	heapStr.Encolar(strs[3])
	heapStr.Encolar(strs[4])

	heapInt.Encolar(nums[0])
	heapInt.Encolar(nums[1])
	heapInt.Encolar(nums[2])
	heapInt.Encolar(nums[3])
	heapInt.Encolar(nums[4])

	require.Equal(t, 5, heapStr.Cantidad())
	require.Equal(t, strs[4], heapStr.VerMax())
	require.False(t, heapStr.EstaVacia())

	require.Equal(t, 5, heapInt.Cantidad())
	require.Equal(t, nums[4], heapInt.VerMax())
	require.False(t, heapInt.EstaVacia())

	heapStr.Desencolar()
	heapInt.Desencolar()

	require.Equal(t, 4, heapStr.Cantidad())
	require.Equal(t, strs[3], heapStr.VerMax())
	require.False(t, heapStr.EstaVacia())

	require.Equal(t, 4, heapInt.Cantidad())
	require.Equal(t, nums[3], heapInt.VerMax())
	require.False(t, heapInt.EstaVacia())

	heapStr.Desencolar()
	heapInt.Desencolar()

	require.Equal(t, 3, heapStr.Cantidad())
	require.Equal(t, strs[2], heapStr.VerMax())
	require.False(t, heapStr.EstaVacia())

	require.Equal(t, 3, heapInt.Cantidad())
	require.Equal(t, nums[2], heapInt.VerMax())
	require.False(t, heapInt.EstaVacia())

	heapStr.Desencolar()
	heapInt.Desencolar()

	require.Equal(t, 2, heapStr.Cantidad())
	require.Equal(t, strs[1], heapStr.VerMax())
	require.False(t, heapStr.EstaVacia())

	require.Equal(t, 2, heapInt.Cantidad())
	require.Equal(t, nums[1], heapInt.VerMax())
	require.False(t, heapInt.EstaVacia())

	heapStr.Desencolar()
	heapInt.Desencolar()

	require.Equal(t, 1, heapStr.Cantidad())
	require.Equal(t, strs[0], heapStr.VerMax())
	require.False(t, heapStr.EstaVacia())

	require.Equal(t, 1, heapInt.Cantidad())
	require.Equal(t, nums[0], heapInt.VerMax())
	require.False(t, heapInt.EstaVacia())

	heapStr.Desencolar()
	heapInt.Desencolar()

	require.Equal(t, 0, heapStr.Cantidad())
	require.True(t, heapStr.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapStr.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapStr.Desencolar() })

	require.Equal(t, 0, heapInt.Cantidad())
	require.True(t, heapInt.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapInt.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapInt.Desencolar() })
}

func TestEncolarYDesencolarRepetidasVeces(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	for i := 0; i < 1000; i++ {
		heap.Encolar(i)
		require.Equal(t, 1, heap.Cantidad())
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
		require.True(t, heap.EstaVacia())
	}

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}

func TestPruebaDeVolumen(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	for i := 1; i <= 10000; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.Cantidad())
		require.Equal(t, i, heap.VerMax())
	}

	require.Equal(t, 10000, heap.Cantidad())
	require.Equal(t, 10000, heap.VerMax())
	require.False(t, heap.EstaVacia())

	i := 10000
	for !heap.EstaVacia() {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
		i--
		require.Equal(t, i, heap.Cantidad())
	}

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSortConNumeros(t *testing.T) {
	entrada := []int{-1, -32, 99, 10, 100, 56}
	salidaEsperada := []int{-32, -1, 10, 56, 99, 100}

	TDAHeap.HeapSort(entrada, func(a, b int) int { return a - b })

	require.Equal(t, entrada, salidaEsperada)
}

func TestHeapSortConStrings(t *testing.T) {
	entrada := []string{"C", "D", "B", "A", "R"}
	salidaEsperada := []string{"A", "B", "C", "D", "R"}

	TDAHeap.HeapSort(entrada, strings.Compare)

	require.Equal(t, entrada, salidaEsperada)
}
