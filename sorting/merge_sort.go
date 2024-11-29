package sorting

func merge(arr *[]int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	// Create temporary arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temporary arrays L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = (*arr)[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = (*arr)[m+1+j]
	}

	// Merge the temporary arrays back into arr[l..r]
	i := 0 // Initial index of first subarray
	j := 0 // Initial index of second subarray
	k := l // Initial index of merged subarray

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			(*arr)[k] = L[i]
			i++
		} else {
			(*arr)[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of L[], if there are any
	for i < n1 {
		(*arr)[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of R[], if there are any
	for j < n2 {
		(*arr)[k] = R[j]
		j++
		k++
	}
}

func MergeSort(arr *[]int, l, r int) {
	if l < r {
		// Same as (l+r)/2, but avoids overflow for large l and r
		m := l + (r-l)/2

		// Sort first and second halves
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)

		merge(arr, l, m, r)
	}
}