package selectionsort

// min
func findSmallInt(arr []int) int {
	min := arr[0]
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > min {
			min = arr[i]
			index = i
		}
	}
	return index
}

// n*n
func selectionSortInt(arr []int) []int {
	len := len(arr)
	newArr := make([]int, len)
	for i := 0; i < len; i++ {
		index := findSmallInt(arr)
		newArr[i] = arr[index]
		arr = append(arr[:index], arr[index+1:]...)
	}
	return newArr
}
