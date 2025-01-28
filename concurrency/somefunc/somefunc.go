package main

import (
	"fmt"
	"time"
)

func SomeFunc(str string) string {
	return fmt.Sprintf("Hello %s", str)
}

func main() {
	// fmt.Println(SomeFunc("German"))
	// fmt.Println(SomeFunc("Bakit"))
	// fmt.Println(SomeFunc("Nikolay"))

	// fmt.Println("Main func is end")

	// go fmt.Println(SomeFunc("German"))
	// go fmt.Println(SomeFunc("Bakit"))
	// go fmt.Println(SomeFunc("Nikolay"))
	// fmt.Println("Main func is end")

	go fmt.Println(SomeFunc("German"))
	go fmt.Println(SomeFunc("Bakit"))
	go fmt.Println(SomeFunc("Nikolay"))
	time.Sleep(2 * time.Millisecond)
	fmt.Println("Main func is end")
}
