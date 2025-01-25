package qeque

import "container/list"

type Deque struct {
	list *list.List
	len  int
}

func New() Deque {
	return Deque{
		list: list.New(),
	}
}

func (d *Deque) Push_back(val any) {
	d.list.PushBack(val)
	d.len++
}

func (d *Deque) Pop_front() any {
	element := d.list.Front()
	d.list.Remove(element)
	d.len--
	return element.Value
}

func (d *Deque) Len() int {
	return d.len
}
