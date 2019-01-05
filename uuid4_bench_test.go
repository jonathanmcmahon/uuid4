package uuid4

import (
	"testing"
)

func BenchmarkNewBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewBytes()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := New()
		if err != nil {
			b.Error(err)
		}
	}
}
