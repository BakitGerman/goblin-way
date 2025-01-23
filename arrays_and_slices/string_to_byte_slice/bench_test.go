package byteslicetostring

import (
	"testing"
	"unsafe"
)

func convert(str string) []byte {
	if len(str) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

var Result []byte

func BenchmarkStdConvert(b *testing.B) {
	str := "Hello! I'm Lion! RRRRRRRRRRRRR!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = []byte(str)
		_ = Result
	}
}

func BenchmarkMyConvert(b *testing.B) {
	str := "Hello! I'm Lion! RRRRRRRRRRRRR!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = convert(str)
		_ = Result
	}
}

func BenchmarkStdConvert2(b *testing.B) {
	str := "Hello! I'm Lion! RRRRRRRRRRRRR!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := []byte(str)
		_ = result
	}
}

func BenchmarkMyConvert2(b *testing.B) {
	str := "Hello! I'm Lion! RRRRRRRRRRRRR!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := convert(str)
		_ = result
	}
}
