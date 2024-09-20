package arrays

func Slices(slice1 []int, slice2 []int) []int {
	// Merge the two slices: a, b, a, c, a, d, b, c, b, d
	
	// Create a new slice to store the merged slices
	
	// A slice is defined by:
	// - a pointer to the first element of the array: ptr (0xc0000b8000)
	// - the length of the slice: len (4)
	// - the capacity of the slice: cap (4)
	mergedSlice := []int{}
	for i := 0; i < len(slice1); i++ {
		mergedSlice = append(mergedSlice, slice1[i])
		mergedSlice = append(mergedSlice, slice2[i])
	}
	return mergedSlice
}