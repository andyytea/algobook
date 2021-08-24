package sort_go

import (
    "fmt"
    "testing"
    "sort"
    "reflect"
)

func testSort (t *testing.T, example []int, sortfn func([]int)) {
    solution_copy := example
    sortfn(example)
    sort.Ints(solution_copy)
    if !reflect.DeepEqual(example, solution_copy) {
        t.Errorf("sorted %v", solution_copy)
        t.Errorf("   got %v", example)
    }
}

func TestSort (t *testing.T) {
    tests := []struct {
        name string
        example []int
    }{
        {"General", []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}},
        {"Empty", []int{}},
        {"Single", []int{1}},
    }
    for _, ex := range tests {
        t.Run(fmt.Sprintf("Range-%s", ex.name), func (t *testing.T) {
            testSort(t, ex.example, InsertionSortRange)
        })
        t.Run(fmt.Sprintf("Indices-%s", ex.name), func (t *testing.T) {
            testSort(t, ex.example, InsertionSortIndices)
        })
    }
}

func benchmarkSort (b *testing.B, size int, sortfn func ([]int)) {
    b.StopTimer()
    for i := 0; i < b.N; i++ {
        data := make([]int, size)
        for i := 0; i < len(data); i++ {
            data[i] = i ^ 0x2cc
        }
        b.StartTimer()
        sortfn(data)
        b.StopTimer()
    }
}

func BenchmarkInsertionSort (b *testing.B) {
    sizeArr := []int{1<<7, 1<<10, 1<<16}
    for _, size := range sizeArr {
        b.Run(fmt.Sprintf("Standard%d", size), func (b *testing.B) {
            benchmarkSort(b, size, sort.Ints)
        })
        b.Run(fmt.Sprintf("Range%d", size), func (b *testing.B) {
            benchmarkSort(b, size, InsertionSortRange)
        })
        b.Run(fmt.Sprintf("Indices%d", size), func (b *testing.B) {
            benchmarkSort(b, size, InsertionSortIndices)
        })
    }
}
