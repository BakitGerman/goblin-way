package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	list := []int{10, 10, 10, 10, 20, 60}
	want := 120
	got := sum(list)
	if want != got {
		t.Errorf("sum(%v), want %d, got %d", list, want, got)
	}
	fmt.Println(max(list))
	fmt.Println(list[binarySearch(list, 0, len(list)-1, 10000)])
	t.Error()

}
