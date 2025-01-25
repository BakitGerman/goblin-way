package main

func main() {
	graphV1()
	graphV2()
	graphV3()
}

func graphV1() [][]int {
	// minus
	// memory overruns
	// there is a maximum possible number of children for a node (there are 8 here)
	// the graph is only undirected
	// ----------------------------

	// -1 missing node
	return [][]int{
		// node node node
		{1, 2, 3},
		{4, 5, -1},
		{-1, 8, 9},
	}
}

func graphV2() [][]int {
	// minus
	// integers only
	// shifting the array when adding
	// iterating through an array when searching for a parent
	// ----------------------------s

	// plus
	// it can be either oriented or undirected.
	// ----------------------------

	return [][]int{
		// parent_node child_node weight
		{1, 2, 3},
		{2, 5, 6},
		{1, 8, 9},
		{3, 5, 10},
	}
}

func graphV3() [][]int {
	// minus
	// integers only
	// shifting the array when adding
	// iterating through an array when searching for a parent

	// plus
	// it can be either oriented or undirected.

	// 0 missing node
	// w - weight
	return [][]int{
		//    from from from
		// to  3    5    9
		// to  3    -2   1
		// to  7    3    1
		//    w     w    w

		//    from from from
		// to  0    1    0
		// to  1    0   0
		// to  0    0    1
		//    w     w    w

		{1, 2, 3},
		{4, 5, -1},
		{-1, 8, 9},
	}
}
