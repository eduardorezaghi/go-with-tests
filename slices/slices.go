package main

import (
	"fmt"
	"slices"
	"math/rand"
)

func RandomWords(n int) []string {
	words := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	result := make([]string, n)

	for i := 0; i < n; i++ {
		result[i] = words[rand.Intn(len(words))]
	}

	return result
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}

	// Creating an empty slice
	var empty []int // nil slice, which needs to be initialized before use
	fmt.Println("nil slice:", empty)

	// empty slice with length 10. The capacity is 10
	empty = make([]int, 0, 10) // type, length, capacity
	fmt.Println("empty slice:", empty)
	// generates panic


	full := append(empty, 1, 2, 3, 4, 5)
	fmt.Println("full (appended) slice:", full)
	// capacity is doubled when the slice is nearly full (about 25% free space)
	fmt.Println("full (appended) slice length & cap:", len(full), cap(full))

	fmt.Println(x)
	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))

	// Generate a slice of random words
	s := RandomWords(5)
	fmt.Println(slices.BinarySearch(s, "The"))
}
