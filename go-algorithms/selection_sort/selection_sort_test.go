package selectionsort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := [...]int{9, -9, 0, -100, 1, 6, 3, 7, 55, 3}
	newArr := selectionSortInt(arr[:])
	testArr := [...]int{55, 9, 7, 6, 3, 3, 1, 0, -9, -100}
	if [10]int(newArr) != testArr {
		t.Errorf("want %v, get %v", testArr, newArr)
	}
}
