package strings

import (
	"strings"
	"testing"
)

func TestUpcase(t *testing.T) {
	got := Upcase("hello")
	want := "HELLO"
	assertStrings(t, got, want)
}

func TestLowercase(t *testing.T) {
	got := Lowercase("HELLO")
	want := "hello"
	assertStrings(t, got, want)
}

func TestMap(t *testing.T) {
	got := Map("hello", func(s string) string {
		return strings.ToUpper(s)
	})
	want := "HELLO"
	assertStrings(t, got, want)
}

func TestClone(t *testing.T) {
	got := Clone("hello")
	want := "hello"
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
