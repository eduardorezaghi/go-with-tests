package integers

import "fmt"

// Adds two integers together and returns the result.
func Add(x, y int) int {
	return x + y
}

// ExampleAdd demonstrates how to use the Add function.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// godoc -http=:6060