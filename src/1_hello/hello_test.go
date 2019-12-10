package main

import "testing"

func TestHello(t *testing.T) {
	singNggenaKoen := func(t *testing.T, got, want string) {
		t.Helper() // supaya kalo ada error ditunjukkin line number yang sesuai, bukan di refer ke method ini
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people's name", func(t *testing.T) {
		got := Hello("Tian")
		want := "Hello, Tian!"
		singNggenaKoen(t, got, want)
	})
	// failing test
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		singNggenaKoen(t, got, want)
	})

	t.Run("say hello in indonesian", func(t *testing.T) {
		got := Hello("Tian", "Indo")
		want := "Halo, Tian"
		singNggenaKoen(t, got, want)
	})
}
