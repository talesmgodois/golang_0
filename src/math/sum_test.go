package math

import "testing"

func TestSum(t *testing.T) {
	got := Sum(2, 3)
	want := 5
	if got != 5 {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSum2(t *testing.T) {
	got := Sum(20, 70)
	want := 90
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
