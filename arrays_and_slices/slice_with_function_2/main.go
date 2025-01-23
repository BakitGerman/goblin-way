package main

import "fmt"

func main() {
	data := make([]int, 3, 6)
	fmt.Println("initial slice: ", data)
	proccess(data)
	fmt.Println("after slice: ", data)
	// fmt.Println("after proccess: ", data[:4])
}

func proccess(data []int) {
	data = append(data, 5)
}
