package recursivesum

func sum(x []int) int {
	if len(x) == 0 {
		return 0
	}
	return x[0] + sum(x[1:])
}
