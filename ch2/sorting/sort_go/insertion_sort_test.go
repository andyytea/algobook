package sort_go

import (
    "testing"
    "sort"
    "reflect"
)


func TestInsertionSort1 (t *testing.T) {
    var intarr = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
    a := intarr
    InsertionSort1(a)
    sorted := intarr
    sort.Ints(sorted)
    if !reflect.DeepEqual(a, sorted) {
        t.Errorf("sorted %v", sorted)
        t.Errorf("   got %v", a)
    }
}

func TestInsertionSort2 (t *testing.T) {
    var intarr = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
    a := intarr
    InsertionSort2(a)
    sorted := intarr
    sort.Ints(sorted)
    if !reflect.DeepEqual(a, sorted) {
        t.Errorf("sorted %v", sorted)
        t.Errorf("   got %v", a)
    }
}

func BenchmarkSortIntStandard128(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<7)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0x2cc
        }
        b.StartTimer()
        sort.Ints(data)
        b.StopTimer()
    }
}

func BenchmarkInsertionSort1_128(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<7)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0x2cc
        }
        b.StartTimer()
        InsertionSort1(data)
        b.StopTimer()
    }
}

func BenchmarkInsertionSort2_128(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<7)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0x2cc
        }
        b.StartTimer()
        InsertionSort2(data)
        b.StopTimer()
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

func BenchmarkInsertionSort1_1K(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<10)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0x2cc
        }
        b.StartTimer()
        InsertionSort1(data)
        b.StopTimer()
    }
}

func BenchmarkInsertionSort2_1K(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<10)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0x2cc
        }
        b.StartTimer()
        InsertionSort2(data)
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

func BenchmarkInsertionSort1_64K(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<16)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0xcccc
        }
        b.StartTimer()
        InsertionSort1(data)
        b.StopTimer()
    }
}

func BenchmarkInsertionSort2_64K(b *testing.B) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, 1<<16)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0xcccc
        }
        b.StartTimer()
        InsertionSort2(data)
        b.StopTimer()
    }
}
