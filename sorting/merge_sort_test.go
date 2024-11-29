package sorting

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/exp/rand"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{12, 11, 13, 5, 6}, []int{5, 6, 11, 12, 13}},
		{[]int{4, 3, 2, 10, 12, 1, 5, 6}, []int{1, 2, 3, 4, 5, 6, 10, 12}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		arr := make([]int, len(test.input))
		copy(arr, test.input)
		MergeSort(&arr, 0, len(arr)-1)
		if !reflect.DeepEqual(arr, test.expected) {
			t.Errorf("got %v, want %v", arr, test.expected)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	// Seed the random number generator for consistent benchmarking
	rand.Seed(uint64(time.Now().UnixNano()))

	// Create a large slice of random integers
	size := 10_000 // Size of the slice
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100000)
	}

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to sort each time
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		MergeSort(&arrCopy, 0, len(arrCopy)-1)
	}
}
