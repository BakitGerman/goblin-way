package main

import (
	"fmt"
	"time"
)

func main() {

	myChannel := make(chan string)

	go func() {
		myChannel <- "Hello "
	}()

	go func() {
		myChannel <- "Bakit"
	}()

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)

	time.Sleep(10 * time.Millisecond)

}
