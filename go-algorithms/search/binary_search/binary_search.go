package binarysearch

// O(log(n)) - worst case
// O(1) - best case
// O(log(n)) - average case
// ω(1) and o(log(n))
func iterativeBinarySearchInts(list []int, target int) int {
	var low int
	high := len(list) - 1
	// moves left or right
	for low <= high {
		mid := int(uint(low+high) >> 1)
		guess := list[mid]
		if guess < target {
			low = mid + 1
		} else if guess > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return low
}

func recursionBinarySearchInts(list []int, target, low, high int) int {
	// based case
	if low > high {
		return low
	}

	mid := int(uint(low+high) >> 1)
	// recursion case
	if list[mid] < target {
		return recursionBinarySearchInts(list, target, mid+1, high)
	} else if list[mid] > target {
		return recursionBinarySearchInts(list, target, low, mid-1)
	}
	return mid
}
