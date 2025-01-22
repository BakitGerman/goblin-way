package binarysearch

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
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
			target:   6,
			expected: 3,
		},
		{
			name:     "Target smaller than all elements",
			list:     []int{1, 3, 5, 7, 9},
			target:   0,
			expected: 0,
		},
		{
			name:     "Target larger than all elements",
			list:     []int{1, 3, 5, 7, 9},
			target:   10,
			expected: 5,
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
			result := binarysearchInt(tt.list, tt.target)
			if result != tt.expected {
				t.Errorf("binarysearchInt(%v, %d) = %d; want %d", tt.list, tt.target, result, tt.expected)
			} else {
				t.Logf("Successfully found number: %d at index %d", tt.target, tt.expected)
			}
		})
	}
}

func generateRandomPhoneNumber() string {
	return fmt.Sprintf("89%09d", rand.Intn(1_000_000_000))
}

func TestBinarySearchString(t *testing.T) {

	rand.NewSource((time.Now().UnixNano()))

	phones := make([]string, 10)
	for i := 0; i < 10; i++ {
		phones[i] = generateRandomPhoneNumber()
	}

	sort.Strings(phones)

	t.Log("Sorted phone numbers:")
	for _, phone := range phones {
		t.Log(phone)
	}

	targets := append(phones, "89878553232", "89999999999")

	for _, target := range targets {
		index := binarySearchString(phones, target)

		if index >= len(phones) || phones[index] != target {
			t.Errorf("Failed to find phone number: %s", target)
		} else {
			t.Logf("Successfully found phone number: %s at index %d", target, index)
		}
	}
}
