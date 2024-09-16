// https://pkg.go.dev/strings#pkg-overview
package strings

import "strings"

// Upcase takes a string and returns it in uppercase
func Upcase(s string) string {
	return strings.ToUpper(s)
}

func Lowercase(s string) string {
	return strings.ToLower(s)
}

func Map(s string, f func(string) string) string {
	return f(s)
}

func Clone(s string) string {
	return strings.Clone(s)
}