package rangebyarraycomparison_test

import "testing"

type account struct {
	balance int
}

func BenchmarkWithPointers(t *testing.B) {
	accounts := [...]*account{{100}, {200}, {300}}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for _, account := range accounts {
			account.balance /= 20
		}
	}
}

func BenchmarkWithoutPointers(t *testing.B) {
	accounts := [...]*account{{100}, {200}, {300}}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for idx := range accounts {
			accounts[idx].balance /= 20
		}
	}
}
