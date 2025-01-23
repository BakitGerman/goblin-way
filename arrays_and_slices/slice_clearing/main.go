package main

import (
	"fmt"
	"slices"
)

func main() {
	// remove all elements
	first := make([]int, 100)
	first = nil
	fmt.Println("first: ", first, " : ", len(first), " : ", cap(first))

	// keep allocated memery
	second := make([]int, 100)
	second = second[:0]
	fmt.Println("first: ", second, " : ", len(second), " : ", cap(second))
	// clip to len(second)
	second = slices.Clip(second)
	fmt.Println("first: ", second, " : ", len(second), " : ", cap(second))
	// zero output elements
	third := make([]int, 100)
	clear(third)
	fmt.Println("first: ", third, " : ", len(third), " : ", cap(third))

}
