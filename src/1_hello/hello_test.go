package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people's name", func(t *testing.T) {
		got := Hello("Tian")
		want := "Hello, Tian!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	// failing test
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
