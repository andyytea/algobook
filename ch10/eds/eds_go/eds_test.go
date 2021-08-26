package eds_go

import (
    "fmt"
    "testing"
)

func TestStack (t *testing.T) {
    // test CreateStackLL
    stk_ll := CreateStackLL()
    stk_arr := CreateStackArr()
    
    var expected_val interface{} = nil
    expected_len := 0
        
    t.Run("create_stackll", func (t *testing.T) {
        got_val := stk_ll.Top()
        got_len := stk_ll.Length()

        if got_val != expected_val {
            t.Errorf("expected '%d' but got '%d'", expected_val, got_val)
        }
        if got_len != expected_len {
            t.Errorf("expected '%d' but got '%d'", expected_len, got_len)
        }
    })

    t.Run("create_stackarr", func (t *testing.T) {
        got_val := stk_arr.Top()
        got_len := stk_arr.Length()

        if got_val != expected_val {
            t.Errorf("expected '%d' but got '%d'", expected_val, got_val)
        }
        if got_len != expected_len {
            t.Errorf("expected '%d' but got '%d'", expected_len, got_len)
        }
    })

    // test Push
    stk_ll.Push(1)
    stk_arr.Push(1)
    expected_val = 1
    expected_len = 1 

    t.Run("push_stackll", func (t *testing.T) {
        got_val := stk_ll.Top().(int)
        got_len := stk_ll.Length()
        if got_val != expected_val {
            t.Errorf("expected '%d' but got '%d'", expected_val, got_val)
        }
        if got_len != expected_len {
            t.Errorf("expected '%d' but got '%d'", expected_len, got_len)
        }
    })
    
    t.Run("push_stackarr", func (t *testing.T) {
        got_val := stk_arr.Top().(int)
        got_len := stk_arr.Length()
        if got_val != expected_val {
            t.Errorf("expected '%d' but got '%d'", expected_val, got_val)
        }
        if got_len != expected_len {
            t.Errorf("expected '%d' but got '%d'", expected_len, got_len)
        }
    })

    expected_len = 0
    
    t.Run("pop_stackll", func (t *testing.T) { 
        popval := stk_ll.Pop().(int)
        got_len := stk_ll.Length()
        if popval != expected_val {
            t.Errorf("expected '%d' but got '%d'", expected_val, popval)
        }
        if got_len != 0 {
            t.Errorf("expected '%d' but got '%d'", expected_len, got_len)
        }
    })

    t.Run("pop_stackarr", func (t *testing.T) { 
        popval := stk_arr.Pop().(int)
        got_len := stk_arr.Length()
        if popval != expected_val {
            t.Errorf("expected '%d' but got '%d'", expected_val, popval)
        }
        if got_len != 0 {
            t.Errorf("expected '%d' but got '%d'", expected_len, got_len)
        }
    })

    t.Run("long_example_stackll", func (t *testing.T) {
        intarr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

        for _, v := range intarr {
            stk_ll.Push(v)
            got := stk_ll.Top().(int)
            want := v
            if got != want {
                t.Errorf("expected '%d' but got '%d'", expected_val, got)
            }
        }
        
        for i := len(intarr) - 1; i >= 0; i-- {
            got := stk_ll.Pop().(int)
            want := intarr[i]
            if got != want {
                t.Errorf("expected '%d' but got '%d'", expected_val, got)
            }
        }
    })

    t.Run("long_example_stackarr", func (t *testing.T) {
        intarr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

        for _, v := range intarr {
            stk_arr.Push(v)
            got := stk_arr.Top().(int)
            want := v
            if got != want {
                t.Errorf("expected '%d' but got '%d'", expected_val, got)
            }
        }
        
        for i := len(intarr) - 1; i >= 0; i-- {
            got := stk_arr.Pop().(int)
            want := intarr[i]
            if got != want {
                t.Errorf("expected '%d' but got '%d'", expected_val, got)
            }
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
