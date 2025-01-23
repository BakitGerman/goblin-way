package main

import (
	"strings"
	"testing"
	"unsafe"
)

func makeDirty(size uint) []byte {
	var sb strings.Builder
	sb.Grow(int(size))
	pointer := unsafe.StringData(sb.String())
	return unsafe.Slice(pointer, size)
}

var Result []byte

func BenchmarkMake(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = make([]byte, 10<<20)
	}

}

func BenchmarkMakeDirty(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = makeDirty(10 << 20)
	}

}
