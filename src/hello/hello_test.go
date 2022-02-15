package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	if len(got) <= 0 {
		t.Errorf("got %s is not expected", got)
	}
}
