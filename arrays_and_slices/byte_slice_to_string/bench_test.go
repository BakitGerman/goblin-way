package byteslicetostring

import (
	"testing"
	"unsafe"
)

func convert(slice []byte) string {
	if len(slice) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(slice), len(slice))
}

var Result string

func BenchmarkStdConvert(b *testing.B) {
	data := make([]byte, 1<<20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = string(data)
		_ = Result
	}
}

func BenchmarkMyConvert(b *testing.B) {
	data := make([]byte, 1<<20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = convert(data)
		_ = Result
	}
}

func BenchmarkStdConvert2(b *testing.B) {
	data := make([]byte, 1<<20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := string(data)
		_ = result
	}
}

func BenchmarkMyConvert2(b *testing.B) {
	data := make([]byte, 1<<20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := convert(data)
		_ = result
	}
}
