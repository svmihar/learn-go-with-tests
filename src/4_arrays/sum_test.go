package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	singNggenaKoen := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("testing adding all 1-5 number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		singNggenaKoen(t, got, want)
	})

	t.Run("testing on slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		singNggenaKoen(t, got, want)
	})
}
func TestSUmAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumTails(t *testing.T) {
	got := SumTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want  %v", got, want)
	}

	got1 := SumTails([]int{}, []int{3, 4, 5})
	want1 := []int{0, 9}
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("got %v want %v", got1, want1)
	}
}
