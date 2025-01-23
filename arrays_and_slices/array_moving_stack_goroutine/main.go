package main

func allocation(index int) int {
	var data [1 << 20]int
	return data[index]
}

func main() {

	array := [100]int{}
	println(&array)

	allocation(1000)

	println(&array)
}
