package main

import (
	"fmt"
	"regexp"
)

// множители. Их эффект будет распространяться на всё, что находится прямо перед ними.
func main() {

	// * ноль или более раз
	regM := regexp.MustCompile(`[1-79]i*`)
	fmt.Println("Example 1")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Ba78it Ge4ma4"))))

	for _, val := range regM.FindAll([]byte("Ba78it Ge4ma4 9"), 4) {
		fmt.Printf("%c, ", val[0])
	}

	// + один или несколько раз
	regM = regexp.MustCompile(`[1-79]i+`)
	fmt.Println("\nExample 2")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Ba78it Ge4ma4 ad6il"))))

	for _, val := range regM.FindAll([]byte("Ba789it Ge4ma4 9"), 4) {
		fmt.Printf("%c, ", val[0])
	}

	// ? один или ноль раз
	regM = regexp.MustCompile(`[1-79]i?`)
	fmt.Println("\nExample 3")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Ba78it Ge4ma4 ad6il"))))

	for _, val := range regM.FindAll([]byte("Ba789it Ge4ma4 9"), 4) {
		fmt.Printf("%c, ", val[0])
	}

	// {n} n - раз
	regM = regexp.MustCompile(`[1-79]i{5}`)
	fmt.Println("\nExample 4")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Ba78it Ge4ma4 ad6il"))))

	for _, val := range regM.FindAll([]byte("Ba789it Ge4ma4 9"), 4) {
		fmt.Printf("%c, ", val[0])
	}

	// {a,b} от а до b раз включительно
	regM = regexp.MustCompile(`[1-79]i{1,3}`)
	fmt.Println("\nExample 5")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Ba78it Ge4ma4 ad6il"))))

	for _, val := range regM.FindAll([]byte("Ba789it Ge4ma4 9"), 4) {
		fmt.Printf("%c, ", val[0])
	}

	// {a} не менее a раз
	regM = regexp.MustCompile(`[1-79]i{2}`)
	fmt.Println("\nExample 5")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Ba78it Ge4ma4 ad6iil"))))

	for _, val := range regM.FindAll([]byte("Ba789it Ge4ma4 9"), 4) {
		fmt.Printf("%c, ", val[0])
	}

	// жадное сопоставление (попытаться найти самую большую строку, которую оно может иметь)
	regM = regexp.MustCompile(`l.*k`)
	fmt.Println("\nExample 6")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Are you looking at the lock or the silk?"))))

	// ленивое сопоставление
	regM = regexp.MustCompile(`l.*?k`)
	fmt.Println("\nExample 6")
	// самое левое совпандение (значение)
	fmt.Println(string(regM.Find([]byte("Are you looking at the lock or the silk?"))))

}
