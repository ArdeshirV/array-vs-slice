package main_test

import (
	"testing"

	"github.com/ArdeshirV/array-vs-slice"
)

func TestArrayAlgorithms(t *testing.T) {
	t.Run("test ArrayAlgorithms", func(t *testing.T) {
		main.ArrayAlgorithms()
	})
}

func TestSliceAlgorithms(t *testing.T) {
	t.Run("test SliceAlgorithms", func(t *testing.T) {
		main.SliceAlgorithms()
	})
}

func BenchmarkArrayAlgorithms(b *testing.B) {
	b.Run("Benchmark ArrayAlgorithms", func(b *testing.B) {
		main.ArrayAlgorithms()
	})
}

func BenchmarkSliceAlgorithms(b *testing.B) {
	b.Run("Benchmark SliceAlgorithms", func(b *testing.B) {
		main.SliceAlgorithms()
	})
}
