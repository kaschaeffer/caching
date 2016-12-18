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

func BenchmarkRandArrayConcurrency1(b *testing.B)  { benchmarkRandArray(b, 1) }
func BenchmarkRandArrayConcurrency2(b *testing.B)  { benchmarkRandArray(b, 2) }
func BenchmarkRandArrayConcurrency8(b *testing.B)  { benchmarkRandArray(b, 8) }
func BenchmarkRandArrayConcurrency16(b *testing.B) { benchmarkRandArray(b, 16) }
func BenchmarkRandArrayConcurrency32(b *testing.B) { benchmarkRandArray(b, 32) }

func BenchmarkSmartRandArrayConcurrency1(b *testing.B)  { benchmarkSmartRandArray(b, 1) }
func BenchmarkSmartRandArrayConcurrency2(b *testing.B)  { benchmarkSmartRandArray(b, 2) }
func BenchmarkSmartRandArrayConcurrency8(b *testing.B)  { benchmarkSmartRandArray(b, 8) }
func BenchmarkSmartRandArrayConcurrency16(b *testing.B) { benchmarkSmartRandArray(b, 16) }
func BenchmarkSmartRandArrayConcurrency32(b *testing.B) { benchmarkSmartRandArray(b, 32) }
