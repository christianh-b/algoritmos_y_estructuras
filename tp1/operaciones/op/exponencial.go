package operaciones

import "math"

func Exponencial(a, b int64) (int64, string) {
	if b < 0 {
		return 0, "ERROR"
	}
	return int64(math.Pow(float64(a), float64(b))), ""
}
