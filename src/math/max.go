package math

func MaxValue(values []float64) float64 {
	var max = 0.0
	for _, s := range values {
		if s > max {
			max = s
		}
	}
	return max
}
