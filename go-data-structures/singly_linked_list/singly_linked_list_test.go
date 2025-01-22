package main

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	test1 := []string{"Front", "Ivan", "Kolya", "Denis", "Vitya", "Arnold"}
	l := LinkedList[string]{}
	l.push_back("Ivan")
	l.push_back("Kolya")
	l.push_back("Denis")
	l.push_back("Vitya")
	l.push_back("Arnold")
	l.push_front("Front")
	output := func(s []string) {
		for i := 0; i < l.size(); i++ {
			val := l.get(i)
			if val != s[i] {
				t.Error(fmt.Errorf("want %s get %s", test1[i], val))
			}
		}
	}
	output(test1)
	l.pop_front()
	test2 := []string{"Ivan", "Kolya", "Denis", "Vitya", "Arnold"}
	output(test2)
	l.pop_back()
	test3 := []string{"Ivan", "Kolya", "Denis", "Vitya", "Arnold"}
	output(test3)
	l.clear()
	if l.size() != 0 {
		t.Error(fmt.Errorf("want %d get %d", 0, l.size()))
	}
	l.push_back("Ivan")
	l.push_back("Kolya")
	l.push_back("Denis")
	l.push_back("Vitya")
	l.push_back("Arnold")
	l.insert("Hoakin Bakli", 3)
	for i := 0; i < l.size(); i++ {
		fmt.Println(l.get(i))
	}
	l.removeAt(3)
	for i := 0; i < l.size(); i++ {
		fmt.Println(l.get(i))
	}
}
