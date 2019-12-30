package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello for peoples", func(t *testing.T) {
		got := Hello("Vitor", "")
		want := "Hello, Vitor"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when a empty string is suplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Vitor", "Spanish")
		want := "Hola, Vitor"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Vitor", "French")
		want := "Bonjour, Vitor"
		assertCorrectMessage(t, got, want)
	})
}
