package math

func Absi(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}
