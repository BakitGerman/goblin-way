package main

import "fmt"

func main() {
	data := make([]int, 3, 6)

	array := (*[3]int)(data[0:3])
	fmt.Println("array: ", array, "data: ", data)

	data[0] = 10

	fmt.Println("array: ", array, "data: ", data)
}
