package insertionsort

import (
	"reflect"
	"testing"
)

var example = []int{13, 7, 6, 45, 21, 9, 101, 102, -23}
var solution = []int{-23, 6, 7, 9, 13, 21, 45, 101, 102}

func TestInsertionSort(t *testing.T) {
	got := InsertionSort(example)
	want := solution

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
