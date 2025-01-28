package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 6)

	// fmt.Println(slice[4]) // panic

	slice = slice[:6]
	fmt.Println(slice[4])
}
