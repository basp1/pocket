package util

func Prob(size int, values []interface{}, probs []float32) []interface{} {
	result := make([]interface{}, 0, size)

	n := len(probs)
	var sum float32 = 0
	for i := 0; i < n; i++ {
		sum += probs[i]
	}
	for i := 0; i < n; i++ {
		p := probs[i] / sum
		from := len(result)
		to := from + int(float32(size)*p)

		for j := from; j < to; j++ {
			result = append(result, values[i])
		}
	}

	return result
}

func ProbInt(size int, values []int, probs []float32) []int {
	result := make([]int, 0, size)

	n := len(probs)
	var sum float32 = 0
	for i := 0; i < n; i++ {
		sum += probs[i]
	}
	for i := 0; i < n; i++ {
		p := probs[i] / sum
		from := len(result)
		to := from + int(float32(size)*p)

		for j := from; j < to; j++ {
			result = append(result, values[i])
		}
	}

	return result
}
