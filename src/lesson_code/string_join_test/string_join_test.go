package string_join_test

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < 100; j++ {
			s = fmt.Sprintf("%v%v", s, j)
		}
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s strings.Builder
		for j := 0; j < 100; j++ {
			s.WriteString(strconv.Itoa(j))
		}
		_ = s.String()
	}
}

func BenchmarkByteBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s bytes.Buffer
		for j := 0; j < 100; j++ {
			s.WriteString(strconv.Itoa(j))
		}
		_ = s.String()
	}
}

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < 100; j++ {
			s += strconv.Itoa(j)
		}
	}
}
