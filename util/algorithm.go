package util

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PopInt(array []int) (int, []int) {
	return array[len(array)-1], array[:len(array)-1]
}

func LastInt(array []int) int {
	return array[len(array)-1]
}
