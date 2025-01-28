package qeque

import "container/list"

type Qeque struct {
	list *list.List
	len  int
}

func New() Qeque {
	return Qeque{
		list: list.New(),
	}
}

func (d *Qeque) Push_back(val any) {
	d.list.PushBack(val)
	d.len++
}

func (d *Qeque) Pop_front() any {
	element := d.list.Front()
	d.list.Remove(element)
	d.len--
	return element.Value
}

func (d *Qeque) Len() int {
	return d.len
}
