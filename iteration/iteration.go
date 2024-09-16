package iteration

import (
	"fmt"
	"strings"
)

// Repeat takes a character and repeats it a number of times
func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
