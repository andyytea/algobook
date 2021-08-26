package eds_go

// Implementation of stack as Linked List
type stack_ll struct {
    length int
    top *node
}

type node struct {
    value interface{}
    prev *node
}

func CreateStackLL() *stack_ll {
    return &stack_ll{0, nil}
}

func (stk *stack_ll) Length() int {
    return stk.length
}

func (stk *stack_ll) Top() interface{} {
    if stk.length == 0 {
        return nil
    }
    return stk.top.value
}

func (stk *stack_ll) Pop() interface{} {
    if stk.length == 0 {
        return nil
    }
    temp := stk.top
    stk.top = stk.top.prev
    stk.length--
    return temp.value
}

func (stk *stack_ll) Push(val interface{}) {
    new_node := &node{val, stk.top}
    stk.length++
    stk.top = new_node
}

// Implementation of stack as slice of array
type stack_arr struct {
    length int
    arr []interface{}
}

func CreateStackArr() *stack_arr {
    return &stack_arr{0, make([]interface{}, 0, 1)}
}

func (stk *stack_arr) Length() int {
    return stk.length
}

func (stk *stack_arr) Top() interface{} {
    if stk.length == 0 {
        return nil
    }
    return stk.arr[stk.length - 1]
}

func (stk *stack_arr) Pop() interface{} {
    if stk.length == 0 {
        return nil
    }
    val := stk.arr[stk.length - 1]
    stk.length--
    stk.arr = stk.arr[:stk.length]
    return val
}

func (stk *stack_arr) Push(val interface{}) {
    stk.arr = append(stk.arr, val)
    stk.length++
}
