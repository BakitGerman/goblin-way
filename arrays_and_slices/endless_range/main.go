package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	for range data {
		data = append(data, 0)
		fmt.Println("iteration")
	}

	// for i := 0; i < len(data); i++ {
	// 	data = append(data, 0)
	// 	fmt.Println("iteration")
	// }
}
