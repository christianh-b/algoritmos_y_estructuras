package operaciones

import (
	"math"
)

func Logaritmo(a, b int64) (int64, string) {
	if a <= 0 || b <= 1 {
		return 0, "ERROR"
	}
	return int64(math.Log(float64(a)) / math.Log(float64(b))), ""
}
