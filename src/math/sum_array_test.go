package math

import (
	"testing"
)

func TestSumArray(t *testing.T) {
	var input = []float64{1, 2, 3}
	const want = 6
	got := SumArray(input)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
