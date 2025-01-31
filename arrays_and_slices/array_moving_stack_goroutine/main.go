package main

func allocation(index int) int {
	var data [1 << 40]int
	return data[index]
}

// check escape
// go run -gcflags='-m' . | grep escape

// go run ./main.go
func main() {
	array := [100]int{}
	println(&array)

	data := allocation(1000)
	_ = data

	println(&array)

}
