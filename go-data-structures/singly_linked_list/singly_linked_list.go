package main

import (
	"fmt"
)

type Node[T any] struct {
	next *Node[T]
	data T
}

type LinkedList[T any] struct {
	head  *Node[T]
	_size int
}

// O(n)
func (this *LinkedList[T]) push_back(value T) {

	if this.head == nil {
		this.head = &Node[T]{data: value}
		this._size++
	} else {
		current := this.head
		for current.next != nil {
			current = current.next
		}
		current.next = &Node[T]{data: value}
		this._size++
	}
}

// O(1)
func (this *LinkedList[T]) push_front(value T) {
	this.head = &Node[T]{data: value, next: this.head}
	this._size++
}

// O(1)
func (this *LinkedList[T]) pop_front() {
	this.head = this.head.next
	this._size--
}

// O(n)
func (this *LinkedList[T]) clear() {
	for this.head != nil {
		this.pop_front()
	}
}

// O(n)
func (this *LinkedList[T]) pop_back() {
	current := this.head
	prev := current
	for current != nil {
		prev = current
		current = current.next
	}
	prev.next = nil
}

// O(n)
func (this *LinkedList[T]) insert(value T, index int) {
	if index == 0 {
		this.push_front(value)
	} else {
		if this._size < index {
			panic(fmt.Sprintf("index %d is out for scope %d", index, this._size))
		}
		current := this.head
		for index != 1 {
			current = current.next
			index--
		}
		current.next = &Node[T]{current.next, value}
		this._size++
	}
}

func (this *LinkedList[T]) removeAt(index int) {
	if index == 0 {
		this.pop_front()
	} else {
		current := this.head
		for index != 1 {
			current = current.next
			index--
		}
		if current.next.next == nil {
			current.next = nil
		} else {
			current.next = current.next.next
		}
	}
	index--
}

func (this *LinkedList[T]) size() int {
	return this._size
}

// O(n)
func (this *LinkedList[T]) get(index int) T {
	if index == 0 {
		return this.head.data
	}
	current := this.head
	for index != 0 {
		current = current.next
		index--
	}
	return current.data
}
