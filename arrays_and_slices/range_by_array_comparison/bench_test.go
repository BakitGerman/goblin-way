package rangebyarraycomparison_test

import "testing"

type account struct {
	balance int
}

func BenchmarkWithPointers(t *testing.B) {
	accounts := make([]account, 10000)
	for i := 0; i < 10000; i++ {
		accounts[i] = account{i + 1}
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for _, account := range accounts {
			account.balance /= 20
		}
	}
}

func BenchmarkWithoutPointers(t *testing.B) {
	accounts := make([]account, 10000)
	for i := 0; i < 10000; i++ {
		accounts[i] = account{i + 1}
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for idx := range accounts {
			accounts[idx].balance /= 20
		}
	}
}
