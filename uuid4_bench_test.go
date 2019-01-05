package uuid4

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := New()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkNewString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewString()
		if err != nil {
			b.Error(err)
		}
	}
}
