package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := make([]int, 3, 3)
	fmt.Println("initial slice: ", data, " : ", unsafe.SliceData(data))
	proccess(data)
	fmt.Println("after slice: ", data, " : ", unsafe.SliceData(data))
}

func proccess(data []int) {
	data = append(data, 5)
	println(unsafe.SliceData(data))
}
