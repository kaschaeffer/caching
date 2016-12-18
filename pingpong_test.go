package caching

import "testing"

const ARRAY_SIZE = 1000000

func benchmarkRandArray(b *testing.B, concurrency int) {
	for n := 0; n < b.N; n++ {
		RandArray(ARRAY_SIZE, concurrency)
	}
}

func benchmarkSmartRandArray(b *testing.B, concurrency int) {
	for n := 0; n < b.N; n++ {
		SmartRandArray(ARRAY_SIZE, concurrency)
	}
}

func BenchmarkRandArray1Core(b *testing.B)  { benchmarkRandArray(b, 1) }
func BenchmarkRandArray2Core(b *testing.B)  { benchmarkRandArray(b, 2) }
func BenchmarkRandArray8Core(b *testing.B)  { benchmarkRandArray(b, 8) }
func BenchmarkRandArray16Core(b *testing.B) { benchmarkRandArray(b, 16) }
func BenchmarkRandArray32Core(b *testing.B) { benchmarkRandArray(b, 32) }

func BenchmarkSmartRandArray1Core(b *testing.B)  { benchmarkSmartRandArray(b, 1) }
func BenchmarkSmartRandArray2Core(b *testing.B)  { benchmarkSmartRandArray(b, 2) }
func BenchmarkSmartRandArray8Core(b *testing.B)  { benchmarkSmartRandArray(b, 8) }
func BenchmarkSmartRandArray16Core(b *testing.B) { benchmarkSmartRandArray(b, 16) }
func BenchmarkSmartRandArray32Core(b *testing.B) { benchmarkSmartRandArray(b, 32) }
