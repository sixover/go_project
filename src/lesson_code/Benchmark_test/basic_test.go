package Benchmark_test

import (
	"bytes"
	"testing"
)

func BenchmarkFirstTest(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
func BenchmarkFirstAnotherTest(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str := ""
		for _, elem := range elems {
			str += elem
		}
	}
	b.StopTimer()
}
