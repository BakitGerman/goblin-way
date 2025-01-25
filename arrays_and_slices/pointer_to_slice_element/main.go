package main

import (
	"fmt"
	"runtime"
)

func allocation() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func FindElement(data []byte, target byte) *byte {
	for i := 0; i < len(data); i++ {
		if data[i] == target {
			return &data[i]
		}
	}
	return nil
}

func main() {
	data := make([]byte, 1<<30)
	pointer := FindElement(data, 0)
	_ = pointer

	allocation()
	runtime.GC()
	allocation()

	runtime.KeepAlive(data)
}
