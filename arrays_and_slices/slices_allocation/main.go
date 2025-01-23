package main

// go build -gcflags='-m' . | grep escape

// <= 64 kb - stack
// > 64 kb - heap
func main() {
	sliceWithStack := make([]byte, 64<<10)
	_ = sliceWithStack

	sliceWithHeap := make([]byte, 64<<10+1)
	_ = sliceWithHeap
}
