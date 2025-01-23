package slicereservation_test

import "testing"

func BenchmarkWithoutReservation(b *testing.B) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 66, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 23, 3, 4, 5, 6, 7, 4, 2, 3, 4, 5, 4, 3, 2, 2, 3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newSlice := make([]int, 0)
		newSlice = append(newSlice, data...)
		_ = newSlice
	}

}

func BenchmarkWithReservation2(b *testing.B) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 66, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 23, 3, 4, 5, 6, 7, 4, 2, 3, 4, 5, 4, 3, 2, 2, 3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newSlice := make([]int, 0, len(data))
		newSlice = append(newSlice, data...)
		_ = newSlice
	}

}

func BenchmarkWithoutReservation2(b *testing.B) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 66, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 23, 3, 4, 5, 6, 7, 4, 2, 3, 4, 5, 4, 3, 2, 2, 3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newSlice := make([]int, 0)
		for idx := range data {
			newSlice = append(newSlice, data[idx])
		}
		_ = newSlice
	}

}

func BenchmarkWithReservation(b *testing.B) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 66, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 23, 3, 4, 5, 6, 7, 4, 2, 3, 4, 5, 4, 3, 2, 2, 3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newSlice := make([]int, 0, len(data))
		for idx := range data {
			newSlice = append(newSlice, data[idx])
		}
		_ = newSlice
	}

}
