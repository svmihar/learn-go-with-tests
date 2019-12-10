package main

import "testing"

func TestRepeat(t *testing.T) {

	singNggenaKoen := func(t *testing.T, got, want string) {
		t.Helper() // supaya kalo ada error ditunjukkin line number yang sesuai, bukan di refer ke method ini
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("testing default 5 times repeating 'a'", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		singNggenaKoen(t, repeated, expected)
	})

	t.Run("testing with parameters for repeating n times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		singNggenaKoen(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
