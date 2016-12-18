package caching

import (
	"fmt"
	"math/rand"
)

func populate(array []int, offset int, step int, done chan<- bool) {
	for i := offset; i < len(array); i += step {
		array[i] = i
	}
	done <- true
}

func smartPopulate(array []int, start int, end int, done chan<- bool) {
	for i := start; i < end; i++ {
		array[i] = i
	}
	done <- true
}

func chunkRange(size int, nChunks int, n int) (int, int) {
	chunkSize := size / nChunks
	start := chunkSize * n
	end := start + chunkSize
	if end > size {
		end = size
	}
	return start, end

}

func SmartRandArray(size int, nThreads int) []int {
	array := make([]int, size)
	done := make(chan bool)

	for n := 0; n < nThreads; n++ {
		start, end := chunkRange(size, nThreads, n)
		go smartPopulate(array, start, end, done)
	}

	for n := 0; n < nThreads; n++ {
		<-done
	}
	return array
}

func RandArray(size int, nThreads int) []int {
	array := make([]int, size)
	done := make(chan bool)

	for n := 0; n < nThreads; n++ {
		go populate(array, n, nThreads, done)
	}

	for n := 0; n < nThreads; n++ {
		<-done
	}
	return array
}

func main() {
	rand.Seed(42)
	array := RandArray(10000000, 16)

	for n := 0; n < 10; n++ {
		fmt.Println(array[n])
	}

}
