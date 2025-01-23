package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3}
	// copy of array
	for range data {
		fmt.Println("iteration")
	}

	// not copy of array
	for range &data {
		fmt.Println("iteration")
	}
	fmt.Println(data)

	// not copy of array
	for range data[:] {
		fmt.Println("iteration")
	}

}
