package main

import (
	"fmt"
	"math/rand"
)

const (
	N = 10000000
)

func main() {
	fmt.Println("Array vs Slice Performance")
	ArrayAlgorithms()
	SliceAlgorithms()
}

func ArrayAlgorithms() int {
	var arr [N]int
	sum := 0
	for i := range N {
		arr[i] = i * i
	}
	for i := N - 1; i >= 0; i-- {
		sum += arr[i]
	}
	for range N {
		j := rand.Intn(N)
		sum ^= arr[j]
	}
	return sum
}

func SliceAlgorithms() int {
	slice := make([]int, N)
	sum := 0
	for i := range N {
		slice[i] = i * i
	}
	for i := N - 1; i >= 0; i-- {
		sum += slice[i]
	}
	for range N {
		j := rand.Intn(N)
		sum ^= slice[j]
	}
	return sum
}
