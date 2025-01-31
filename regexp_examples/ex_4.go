package main

import (
	"fmt"
	"regexp"
)

// Классы сокращенных символов
func main() {

	// \ экранизация, допустим $ это метасимвол, но с экранизацией это обычный символ
	regM := regexp.MustCompile(`\$\d{4}`)
	fmt.Println("Example 1")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Сегодня я заработал $58327, но потерял $3826."))))

	for _, val := range regM.FindAll([]byte("Сегодня я заработал $58327, но потерял $3826."), 4) {
		fmt.Printf("%s, ", val)
	}

	regM.FindAllString()
}
