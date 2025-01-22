package recursivesum

import "testing"

func TestRecursiveSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 15
	get := sum(arr)
	if want != get {
		t.Errorf("want %d, get %d", want, get)
	}
}
