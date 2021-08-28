package eds_go

import (
    "fmt"
    "testing"
)

func testFn(t *testing.T, got interface{}, expected interface{}) {
    if got != expected {
            t.Errorf("expected '%d' but got '%d'", expected, got)
    }
}

func TestStack (t *testing.T) {
    // test CreateStackLL
    stk_ll := CreateStackLL()
    stk_arr := CreateStackArr()
            
    t.Run("create_stackll", func (t *testing.T) {
        testFn(t, stk_ll.IsEmpty(), true)
        testFn(t, stk_ll.Top(), nil)
        testFn(t, stk_ll.Length(), 0)
    })

    t.Run("create_stackarr", func (t *testing.T) {
        testFn(t, stk_arr.IsEmpty(), true)
        testFn(t, stk_arr.Top(), nil)
        testFn(t, stk_arr.Length(), 0)
    })

    // test Push
    stk_ll.Push(1)
    stk_arr.Push(1)
     

    t.Run("push_stackll", func (t *testing.T) {
        testFn(t, stk_ll.IsEmpty(), false)
        testFn(t, stk_ll.Top(), 1)
        testFn(t, stk_ll.Length(), 1)
    })
    
    t.Run("push_stackarr", func (t *testing.T) {
        testFn(t, stk_arr.IsEmpty(), false)
        testFn(t, stk_arr.Top(), 1)
        testFn(t, stk_arr.Length(), 1)
    }) 
    
    t.Run("pop_stackll", func (t *testing.T) { 
        testFn(t, stk_ll.Pop(), 1)
        testFn(t, stk_ll.Length(), 0)
    })

    t.Run("pop_stackarr", func (t *testing.T) { 
        testFn(t, stk_arr.Pop(), 1)
        testFn(t, stk_arr.Length(), 0)
    })

    t.Run("long_example_stackll", func (t *testing.T) {
        intarr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

        for _, expected := range intarr {
            stk_ll.Push(expected)
            testFn(t, stk_ll.Top(), expected)
        }
        
        for i := len(intarr) - 1; i >= 0; i-- {
            testFn(t, stk_ll.Pop(), intarr[i])
        }
    })

    t.Run("long_example_stackarr", func (t *testing.T) {
        intarr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

        for _, expected := range intarr {
            stk_arr.Push(expected)
            testFn(t, stk_arr.Top(), expected)
        }
        
        for i := len(intarr) - 1; i >= 0; i-- {
            testFn(t, stk_arr.Pop(), intarr[i])
        }
    })
}

func benchmarkPushLL (b *testing.B, size int) {
    stk := CreateStackLL()
    for i := 0; i < size; i++ {
        stk.Push(i ^ 0x2cc)
    } 
}

func benchmarkPushArr (b *testing.B, size int) {
    stk := CreateStackArr()
    for i := 0; i < size; i++ {
        stk.Push(i ^ 0x2cc)
    } 
}

func benchmarkPopLL (b *testing.B, size int) {
    b.StopTimer()
    stk := CreateStackLL()
    for i := 0; i < size; i++ {
        stk.Push(i ^ 0x2cc)
    }
    b.StartTimer()
    for stk.Length() > 0 {
        stk.Pop()
    }
    b.StopTimer()
}

func benchmarkPopArr (b *testing.B, size int) {
    b.StopTimer()
    stk := CreateStackArr()
    b.Log("\n")
    for i := 0; i < size; i++ {
        if i%200000 == 0 {
            b.Logf("cap: %d\n", cap(stk.arr))
            b.Logf("len: %d\n", len(stk.arr))
        }
        stk.Push(i ^ 0x2cc)
    }
    b.StartTimer()
    for stk.Length() > 0 {
        stk.Pop()
    }
    b.StopTimer()
}

func BenchmarkStack (b *testing.B) {
    sizeArr := []int{1<<20}
    for _, size := range sizeArr {
        sizeMil := size/1000
        b.Run(fmt.Sprintf("push_stack_ll_%dK", sizeMil), func (b *testing.B) { benchmarkPushLL(b, size) })
        b.Run(fmt.Sprintf("pop_stack_ll_%dK", sizeMil), func (b *testing.B) { benchmarkPopLL(b, size) })
        b.Run(fmt.Sprintf("push_stack_arr_%dK", sizeMil), func (b *testing.B) { benchmarkPushArr(b, size) })
        b.Run(fmt.Sprintf("pop_stack_arr_%dK", sizeMil), func (b *testing.B) { benchmarkPopArr(b, size) })
    }
}
