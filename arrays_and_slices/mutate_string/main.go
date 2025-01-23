package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := []byte{'g', 'g', 'l', 'i', 'o', 'n'}
	pointer := (*byte)(unsafe.Pointer(unsafe.SliceData(data)))
	str := unsafe.String(pointer, len(data))
	fmt.Printf("initial %s\n", str)
	data[1] = 'o'
	fmt.Printf("modifed %s", str)
}
