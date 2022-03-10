package math

func SumArray(values []float64) float64 {
	var sum = 0.0
	for _, s := range values {
		sum += s
	}
	return sum
}
