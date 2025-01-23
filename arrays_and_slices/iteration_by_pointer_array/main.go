package main

import (
	"fmt"
	"unsafe"
)

func main() {

	data := make([]int, 10, 20)
	println(unsafe.SliceData(data))
	fmt.Println(len(data), cap(data))

	data = data[1:]
	fmt.Println(len(data), cap(data))
	println(unsafe.SliceData(data))

}
