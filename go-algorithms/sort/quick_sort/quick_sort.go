package quicksort

import (
	"math/rand"
	"time"
)

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		index := r.Intn(len(list))
		pivot := list[index]
		var less = []int{}
		var greater = []int{}
		for _, num := range append(list[:index], list[index+1:]...) {
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		less = append(quicksort(less), pivot)
		greater = quicksort(greater)
		return append(less, greater...)
	}
}
