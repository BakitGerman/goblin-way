package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		// Тест 1: Пустой массив
		{
			input:    []int{},
			expected: []int{},
		},
		// Тест 2: Массив с одним элементом
		{
			input:    []int{1},
			expected: []int{1},
		},
		// Тест 3: Уже отсортированный массив
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		// Тест 4: Массив в обратном порядке
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		// Тест 5: Массив с повторяющимися элементами
		{
			input:    []int{4, 2, 4, 1, 3, 2},
			expected: []int{1, 2, 2, 3, 4, 4},
		},
		// Тест 6: Случайный порядок
		{
			input:    []int{10, 7, 8, 9, 1, 5},
			expected: []int{1, 5, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		result := quicksort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Для входного массива %v ожидается %v, но получено %v", test.input, test.expected, result)
		}
	}
}
