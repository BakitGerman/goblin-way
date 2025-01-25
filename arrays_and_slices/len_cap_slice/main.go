package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := make([]int, 1<<5) // 32 elements
	fmt.Println("data :", unsafe.SliceData(data), " : ", len(data), " : ", cap(data))
	fmt.Println("data :", unsafe.SliceData(data[1:]), " : ", len(data[1:]), " : ", cap(data[1:]))
	fmt.Println("data :", unsafe.SliceData(data[:31]), " : ", len(data[:31]), " : ", cap(data[:31]))
}
