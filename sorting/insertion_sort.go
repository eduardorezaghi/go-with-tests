package sorting

import (
	"fmt"
)

// InsertionSort sorts an array using insertion sort algorithm.
func InsertionSort(arr []int) {
	// We start from the second element because we presume the first element is sorted.
	for i := 1; i < len(arr); i++ {
		// Analyze the current element
		key := arr[i]
		j := i - 1
		// Only move elements of arr[0..i-1] that are greater
		// than key to one position ahead of their current position.
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}


func Isort(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i];
		j := i - 1;

		for j >= 0 && arr[j] > value {
			temp := arr[j]
			arr[j+1] = temp
			j--
		}
		arr[j+1] = value
	}
}

func main() {
	// Example usage
	arr := []int{12, 11, 13, 5, 6}
	InsertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
