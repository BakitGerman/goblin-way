package factorial

import "testing"

func TestFactorial(t *testing.T) {
	want := 120
	inpVal := 5
	got := fact(inpVal)
	if got != want {
		t.Errorf("factorial(%d), want %d, got %d", inpVal, want, got)
	}
}
