package sort_go

import (
    "testing"
    "sort"
    "reflect"
)

var intarr = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestInsertionSortN (t *testing.T) {
    a := InsertionSortN(intarr)
    sorted := intarr
    sort.Ints(sorted)
    if !reflect.DeepEqual(a, sorted) {
        t.Errorf("sorted %v", sorted)
        t.Errorf("   got %v", a)
    }
}

func BenchmarkSortIntStandard1K(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<10)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0x2cc
        }
        b.StartTimer()
        sort.Ints(data)
        b.StopTimer()
    }
}

func BenchmarkInsertionSortN1K(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<10)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0x2cc
        }
        b.StartTimer()
        InsertionSortN(data)
        b.StopTimer()
    }
}

func BenchmarkSortIntStandard64K(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<16)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0xcccc
        }
        b.StartTimer()
        sort.Ints(data)
        b.StopTimer()
    }
}

func BenchmarkInsertionSortN64K(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<16)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0xcccc
        }
        b.StartTimer()
        InsertionSortN(data)
        b.StopTimer()
    }
}
