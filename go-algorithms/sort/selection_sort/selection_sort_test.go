package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSortInt(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected []int
	}{
		{
			name:     "Already sorted list",
			list:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted list",
			list:     []int{5, 3, 6, 2, 10},
			expected: []int{2, 3, 5, 6, 10},
		},
		{
			name:     "List with duplicates",
			list:     []int{5, 3, 5, 2, 10},
			expected: []int{2, 3, 5, 5, 10},
		},
		{
			name:     "List with negative numbers",
			list:     []int{-3, 0, -1, 4, 2},
			expected: []int{-3, -1, 0, 2, 4},
		},
		{
			name:     "Empty list",
			list:     []int{},
			expected: []int{},
		},
		{
			name:     "Single element list",
			list:     []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := selectionSortInt(tt.list)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SelectionSortInt(%v) = %v; want %v", tt.list, result, tt.expected)
			} else {
				t.Logf("Successfully sorted: %v", tt.expected)
			}
		})
	}
}
