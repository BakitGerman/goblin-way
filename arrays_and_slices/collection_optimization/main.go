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

func main() {
	data := make([]byte, 1<<30)

	allocation()
	runtime.GC()
	allocation()

	runtime.KeepAlive(len(data))
	fmt.Println(len(data))
}
