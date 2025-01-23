package main

import (
	"fmt"
	"sort"
	"unsafe"
)

type A struct {
	A *int
	B *int
}

// check pointers
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := sort.SearchInts(s, -1)
	fmt.Println(i)

	var a int
	pa := unsafe.Pointer(&a)
	fmt.Println(&a, pa)

	var b int
	A := A{A: &a, B: &b}
	pA := unsafe.Pointer(&A)
	println("Struct A ", &A, pA)

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	parr := unsafe.Pointer(&arr)
	println(&arr, parr, *(*int)(unsafe.Add(parr, 4*unsafe.Sizeof(int(0)))))

	slice := make([]int, 10)
	slice = append(slice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	pslice := unsafe.Pointer(&slice)
	println(&slice, pslice, *(*int)(unsafe.Add(pslice, 4*unsafe.Sizeof(int(0)))))

}
