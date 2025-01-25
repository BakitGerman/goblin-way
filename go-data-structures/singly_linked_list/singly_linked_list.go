package singlylinkedlist

import (
	"fmt"
)

type node[T any] struct {
	next *node[T]
	data T
}

type List[T any] struct {
	head  *node[T]
	_size int
}

// O(n) - worst
// O(1) - best
func (this *List[T]) push_back(value T) {

	if this.head == nil {
		this.head = &node[T]{data: value, next: nil}
		this._size++
	} else {
		current := this.head
		for current.next != nil {
			current = current.next
		}
		current.next = &node[T]{data: value}
		this._size++
	}
}

// O(1) - worst
func (this *List[T]) push_front(value T) {
	this.head = &node[T]{data: value, next: this.head}
	this._size++
}

// O(1)
func (this *List[T]) pop_front() T {
	if this.head == nil {
		panic("attempt to pop_front from an empty list")
	}
	ret := this.head
	this.head = this.head.next
	this._size--
	return ret.data
}

// O(n) - worst
// O(1) - best
func (this *List[T]) clear() {
	this.head = nil
	this._size = 0
}

// O(n)
func (this *List[T]) pop_back() T {
	if this.head == nil {
		panic("attempt to pop_back from an empty list")
	}
	if this.head.next == nil {
		ret := this.head.data
		this.head = nil
		this._size--
		return ret
	}
	current := this.head
	for current.next.next != nil {
		current = current.next
	}
	ret := current.next.data
	current.next = nil
	this._size--
	return ret
}

// O(n) - worst
// O(1) - best
func (this *List[T]) insert(value T, index int) {
	if index < 0 || index > this._size {
		panic(fmt.Sprintf("index %d is out of range [0, %d]", index, this._size))
	}
	if index == 0 {
		this.push_front(value)
		return
	}
	if index == this._size {
		this.push_back(value)
		return
	}
	current := this.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	current.next = &node[T]{data: value, next: current.next}
	this._size++
}

// O(n)
func (this *List[T]) removeAt(index int) {
	if index < 0 || index >= this._size {
		panic(fmt.Sprintf("index %d is out of range [0, %d)", index, this._size))
	}
	if index == 0 {
		this.pop_front()
		return
	}
	if index == this._size-1 {
		this.pop_back()
		return
	}
	current := this.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	current.next = current.next.next
	this._size--
}

func (this *List[T]) size() int {
	return this._size
}

// O(n)
func (this *List[T]) get(index int) T {
	if index < 0 || index >= this._size {
		panic(fmt.Sprintf("index %d is out of range [0, %d)", index, this._size))
	}
	current := this.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data
}
