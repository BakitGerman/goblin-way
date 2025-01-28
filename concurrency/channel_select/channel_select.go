package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	myChannel := make(chan string)
	otherChannel := make(chan string)

	go func() {
		myChannel <- "Hello "
	}()

	strings.Split("sdfsdfs dsf", " ")

	go func() {
		otherChannel <- "Bakit"
	}()

	select {
	case msg := <-myChannel:
		fmt.Println(msg)
	case msg := <-otherChannel:
		fmt.Println(msg)
	}

	time.Sleep(10 * time.Millisecond)
}
