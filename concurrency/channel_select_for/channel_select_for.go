package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)
	otherChannel := make(chan string)

	go func() {
		myChannel <- "Hello "
		myChannel <- "Hello "
		myChannel <- "Hello "
		myChannel <- "Hello "
		myChannel <- "Hello "
	}()

	go func() {
		otherChannel <- "Bakit"
	}()

	for {
		select {
		case msg := <-myChannel:
			fmt.Println(msg)
		case msg := <-otherChannel:
			fmt.Println(msg)
		}
	}

	// time.Sleep(10 * time.Millisecond)
}
