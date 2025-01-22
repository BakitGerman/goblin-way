package sum

func sum(list []int) int {
	if len(list) == 0 {
		return 0
	}
	if len(list) == 1 {
		return list[0]
	}
	return list[0] + sum(list[1:])
}

func max(list []int) int {
	if len(list) == 0 {
		return 0
	}
	if len(list) == 1 {
		return list[0]
	}
	val := list[0]
	nextVal := max(list[1:])
	if val > nextVal {
		return val
	} else {
		return nextVal
	}
}

func binarySearch(list []int, low, high, target int) int {
	if low >= high {
		return low
	}
	mid := int(uint(low+high) >> 1)
	if list[mid] == target {
		return mid
	} else if list[mid] < target {
		return binarySearch(list, mid+1, high, target)
	} else if list[mid] > target {
		return binarySearch(list, low, mid-1, target)
	}
	return low
}
