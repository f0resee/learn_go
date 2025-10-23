package testing

import (
	"fmt"
	"testing"
)

func Test_T(t *testing.T) {
	t.Run("T", func(t *testing.T) {
		fmt.Printf("T\n")
	})
}

func Benchmark_B(b *testing.B) {
	b.Run("B", func(b *testing.B) {
		fmt.Printf("B\n")
	})
}
