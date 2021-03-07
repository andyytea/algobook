package selectionsort

import (
	"reflect"
	"testing"
)

var empty = []int{}
var example = []int{13, 7, 6, 45, 21, 9, 101, 102, -23}
var solution = []int{-23, 6, 7, 9, 13, 21, 45, 101, 102}

func TestSelectionSort(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("empty sorting case", func(t *testing.T) {
		got := SelectionSort(empty)
		want := empty
		assertCorrectMessage(t, got, want)
	})

	t.Run("general sorting case", func(t *testing.T) {
		got := SelectionSort(example)
		want := solution
		assertCorrectMessage(t, got, want)
	})
}
