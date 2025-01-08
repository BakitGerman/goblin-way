package binarysearch

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestOKBinarySearchInt(t *testing.T) {
	var vals [100]int
	for i := 0; i < 100; i++ {
		vals[i] = i + 1
	}
	targets := [...]int{1, 5, 55, 33, 100, 93, 92}
	println("my binary search")
	println("========================================================")
	for i := 0; i < len(targets); i++ {
		index := binarysearchInt(vals[:], targets[i])
		if targets[i] != vals[index] {
			t.Errorf("we want an number %d, received %d", targets[i], vals[index])
		}
	}
	println("std binary search")
	println("========================================================")
	// std binary search
	for j := 0; j < len(targets); j++ {
		index := sort.Search(len(vals), func(i int) bool {
			return vals[i] >= targets[j]
		})
		if targets[j] != vals[index] {
			t.Errorf("we want an number %d, received %d", targets[j], vals[index])
		}
	}
}

func TestFailBinarySearchInt(t *testing.T) {
	var vals [100]int
	for i := 0; i < 100; i++ {
		vals[i] = i + 1
	}
	targets := [...]int{-1, 0, 101}
	println("my binary search")
	println("========================================================")
	for i := 0; i < len(targets); i++ {
		index := binarysearchInt(vals[:], targets[i])
		if index >= len(vals) {
			t.Errorf("number %d not find", targets[i])
		} else if targets[i] != vals[index] {
			t.Errorf("we want an number %d, received %d", targets[i], vals[index])
		}
	}
}

func generateRandomPhoneNumber() string {
	return fmt.Sprintf("89%09d", rand.Intn(1_000_000_000))
}

func TestBinarySearchString(t *testing.T) {
	rand.NewSource(time.Now().Unix())
	phones := make([]string, 10)
	for i := 0; i < 10; i++ {
		phones[i] = generateRandomPhoneNumber()
	}
	sort.Strings(phones)
	println("Sorting phone numbers:")
	for _, phone := range phones {
		println(phone)
	}
	println("Find...")
	targets := phones
	targets = append(targets, "89878553232", "89999999999")
	for j := 0; j < len(targets); j++ {
		index := binarySearch(len(phones), func(i int) bool {
			return phones[i] >= targets[j]
		})
		if index >= len(phones) {
			t.Errorf("number %s cannot be find", targets[j])
		} else if targets[j] != phones[index] {
			t.Errorf("we want an number %s, received %s", targets[j], phones[index])
		}
	}
}
