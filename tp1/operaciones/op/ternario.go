package operaciones

func OperadorTernario(a, b, c int64) int64 {
	if a != 0 {
		return b
	}
	return c
}
