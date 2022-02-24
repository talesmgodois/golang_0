package math

import (
	"testing"
)

func TestMax(t *testing.T) {
	var input = []float64{10.0, 5.0, 3.4, 90}
	const want = 90
	got := MaxValue(input)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
