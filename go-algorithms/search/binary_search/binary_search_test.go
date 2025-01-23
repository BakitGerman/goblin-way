package binarysearch

import (
	"testing"
)

func TestBinarySearchInt(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		target   int
		expected int
	}{
		{
			name:     "Target in list",
			list:     []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 2, // 5 находится на индексе 2
		},
		{
			name:     "Target not in list, insert position",
			list:     []int{1, 3, 5, 7, 9},
			target:   10,
			expected: 5,
		},
		{
			name:     "Target smaller than all elements",
			list:     []int{1, 3, 5, 7, 9},
			target:   0,
			expected: 0,
		},
		{
			name:     "Target larger than all elements",
			list:     []int{1, 3, 5, 5, 7, 9},
			target:   5,
			expected: 2,
		},
		{
			name:     "Empty list",
			list:     []int{},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element matching",
			list:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element not matching",
			list:     []int{5},
			target:   3,
			expected: 0,
		},
		{
			name:     "Duplicate elements in list",
			list:     []int{1, 3, 3, 5, 7},
			target:   3,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := iterativeBinarySearchInts(tt.list, tt.target)
			if result != tt.expected {
				t.Errorf("binarysearchInt(%v, %d) = %d; want %d", tt.list, tt.target, result, tt.expected)
			} else {
				t.Logf("Successfully found number: %d at index %d", tt.target, tt.expected)
			}
		})
	}
	// recursion binary search tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := recursionBinarySearchInts(tt.list, tt.target, 0, len(tt.list)-1)
			if result != tt.expected {
				t.Errorf("binarysearchInt(%v, %d) = %d; want %d", tt.list, tt.target, result, tt.expected)
			} else {
				t.Logf("Successfully found number: %d at index %d", tt.target, tt.expected)
			}
		})
	}
}
