package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlices(t *testing.T) {
	myRuneSlice := []int{'a', 'b', 'c', 'd'}
	myIntSlice := []int{1, 2, 3, 4}
	var mergedSlice = Slices(myRuneSlice, myIntSlice)
	want := []int{'a', 1, 'b', 2, 'c', 3, 'd', 4}

	assert.Equal(t, want, mergedSlice)
}
