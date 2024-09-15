package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Edward!", "")
		want := "Hello, Edward!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, Go!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello, Go! in Spanish", func(t *testing.T) {
		got := Hello("Go!", "Spanish")
		want := "Hola, Go!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello, Go! in French", func(t *testing.T) {
		got := Hello("Go!", "French")
		want := "Bonjour, Go!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
