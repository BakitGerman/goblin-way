package main

import "fmt"

func main() {
	array := [10]int{}

	for idx, val := range array {
		val += 50
		println(&val, &array[idx])
	}
	fmt.Println(array)
}
