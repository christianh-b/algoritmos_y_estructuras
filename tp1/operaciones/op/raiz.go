package operaciones

import "math"

func Raiz(a int64) (int64, string) {
	if a < 0 {
		return 0, "ERROR"
	}
	return int64(math.Sqrt(float64(a))), ""
}
