package test

import (
	"testing"
	"github.com/KseniiaL/Making-tree-in-go/tree"
)

//benchmark test for creating trees
func benchmarkNew(size int, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.New(size, 1)
	}
}

//benchmark test for inserting values into trees
func benchmarkInsert(size int, b *testing.B) {
	b.ResetTimer()
	t := tree.New(size, 1)
	for n := 0; n < b.N; n++ {
		tree.Insert(t, 1)
	}
}

func BenchmarkNew10k(b *testing.B) { benchmarkNew(10000, b) }
func BenchmarkNew20k(b *testing.B) { benchmarkNew(20000, b) }
func BenchmarkNew30k(b *testing.B) { benchmarkNew(30000, b) }
func BenchmarkNew40k(b *testing.B) { benchmarkNew(40000, b) }
func BenchmarkNew50k(b *testing.B) { benchmarkNew(50000, b) }

func BenchmarkInsert10k(b *testing.B) { benchmarkInsert(10000, b)}
func BenchmarkInsert20k(b *testing.B) { benchmarkInsert(20000, b)}
func BenchmarkInsert30k(b *testing.B) { benchmarkInsert(30000, b)}
func BenchmarkInsert40k(b *testing.B) { benchmarkInsert(40000, b)}
func BenchmarkInsert50k(b *testing.B) { benchmarkInsert(50000, b)}