package doublelinkedlist

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	value T
}

type List[T any] struct {
	head *Node[T]
	len  int
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (this *List[T]) push_back(value T) {
	if this.head == nil {
		this.head = &Node[T]{value: value}
	}

}
