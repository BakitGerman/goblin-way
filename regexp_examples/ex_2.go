package main

import (
	"fmt"
	"regexp"
)

func main() {

	regM := regexp.MustCompile(`[1-79]`)
	fmt.Println("Example 1")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Ba78it Ge4ma4"))))

	for _, val := range regM.FindAll([]byte("Ba78it Ge4ma4 9"), 4) {
		fmt.Printf("%c, ", val[0])
	}

	fmt.Println("\nExample 2")
	regM = regexp.MustCompile(`[^1-79]`)
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Ba78it Ge4ma4"))))

	for _, val := range regM.FindAll([]byte("Ba78it Ge4ma4 9"), len("Ba78it Ge4ma4 9")) {
		fmt.Printf("%c, ", val[0])
	}

}
