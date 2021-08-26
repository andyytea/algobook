# Naive Implementation of stack as linked list and as array
```
goos: darwin
goarch: arm64
pkg: algobook/ch10/eds/eds_go
BenchmarkStack/push_stack_ll_1048K-8         	1000000000	         0.05872 ns/op	       0 B/op	       0 allocs/op
BenchmarkStack/pop_stack_ll_1048K-8          	1000000000	         0.003018 ns/op	       0 B/op	       0 allocs/op
BenchmarkStack/push_stack_arr_1048K-8        	1000000000	         0.03870 ns/op	       0 B/op	       0 allocs/op
BenchmarkStack/pop_stack_arr_1048K-8         	1000000000	         0.002202 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	algobook/ch10/eds/eds_go	1.704s
```

Honestly, I am quite confused about how the benchmark is measuring allocations. I would have thought that both implementations involve memory escaping to the heap in the benchmark. In the case of the array, the initial creation of the stack initializes the field `arr` to a slice with length 0, capacity 1. Thus, it would have to grow with every push value. Observe the following escape analysis:

```
go build -gcflags=-m

# algobook/ch10/eds/eds_go
...
./eds.go:15:12: &stack_ll{...} escapes to heap
./eds.go:18:7: stk does not escape
...
./eds.go:40:17: &node{...} escapes to heap
./eds.go:52:12: &stack_arr{...} escapes to heap
./eds.go:52:30: make([]interface {}, 0, 1) escapes to heap
./eds.go:55:7: stk does not escape
```

Additionally, with some scuffed logging we also see that the slice capacity does in fact grow as more values are appended onto the underlying array. Observe:

```
eds_test.go:168: cap: 1
eds_test.go:169: len: 0
eds_test.go:168: cap: 232448
eds_test.go:169: len: 200000
eds_test.go:168: cap: 454656
eds_test.go:169: len: 400000
eds_test.go:168: cap: 710656
eds_test.go:169: len: 600000
eds_test.go:168: cap: 888320
eds_test.go:169: len: 800000
eds_test.go:168: cap: 1110528
eds_test.go:169: len: 1000000
```

Chances are, I'm doing something horribly wrong here. Note that the capacity should represent the size of the underlying array. What I expected is that Go has some arbitrary method for growing the size of the underlying array and minimizing the number of `realloc` operations needed. However, the benchmark suggests that no heap allocations were made at all, how is this possible? I tried to find an answer, but I've given up at this point. I'll come back to this some day.
