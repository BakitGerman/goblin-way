package binarysearch

func binarysearchInt(list []int, target int) int {
	high := len(list) - 1
	low := 0
	for low <= high {
		mid := int(uint(low+high) >> 1)
		guess := list[mid]
		if guess == target {
			return mid
		} else if guess < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func binarySearchString(list []string, target string) int {
	return binarySearch(len(list), func(i int) bool {
		return list[i] >= target
	})
}

func binarySearch(n int, f func(i int) bool) int {
	low, high := 0, n-1
	for low <= high {
		mid := int(uint(low+high) >> 1)
		if !f(mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
