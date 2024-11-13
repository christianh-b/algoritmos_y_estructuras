package operaciones

func Division(a, b int64) (int64, string) {
	if b == 0 {
		return 0, "ERROR"
	}
	return a / b, ""
}
