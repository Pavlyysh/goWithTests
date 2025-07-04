// 1. Write a test
// 2. Make the compiler pass
// 3. Run the test, see that it fails and check the error message is meaningful
// 4. Write enough code to make the test pass
// 5. Refactor

package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Pablo", "Spanish")
		want := "Hola, Pablo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Paul", "French")
		want := "Bonjour, Paul"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Pasha", "Russian")
		want := "Privet, Pasha"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
