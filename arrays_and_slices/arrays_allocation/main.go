package main

// go build -gcflags='-m' . | grep escape

func allocation() *[10]int {
	var data [10]int
	return &data
}

func main() {
	var array [1 << 20]int8
	_ = array

	var array2 [1<<40 + 1]int8
	_ = array2

	allocation()
}
