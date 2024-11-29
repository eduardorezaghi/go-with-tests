package main


import (
	"fmt"
	"maps"
)

func main() {
	myMap := make(map[rune]string)

	myMap['a'] = "apple"
	myMap['b'] = "banana"
	myMap['c'] = "cherry"

	fmt.Println(myMap)

	clonedMap := maps.Clone(myMap)
	clonedMap['d'] = "date"
	fmt.Println(clonedMap)
}