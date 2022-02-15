package uuid

import "testing"

func TestRandom(t *testing.T) {
	got, err := Uuid()

	if err != nil {
		t.Errorf("Uuid Not generated %s", got)
	}
	if &got == nil {
		t.Errorf("Uuid could not be generated %s", got)
	}
}
