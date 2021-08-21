# Naive implementations of sorting in Go vs. Standard Library

```
goos: darwin
goarch: arm64
pkg: algobook/ch2/sorting/sort_go
BenchmarkSortIntStandard128-8   	  274051	      3943 ns/op	      24 B/op	       1 allocs/op
BenchmarkInsertionSort1_128-8   	  232780	      5157 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort2_128-8   	  229870	      5185 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortIntStandard1K-8    	   25922	     45202 ns/op	      24 B/op	       1 allocs/op
BenchmarkInsertionSort1_1K-8    	    4932	    242133 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort2_1K-8    	    4909	    244160 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortIntStandard64K-8   	     230	   5204274 ns/op	      24 B/op	       1 allocs/op
BenchmarkInsertionSort1_64K-8   	       2	 608789354 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort2_64K-8   	       2	 616978625 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	algobook/ch2/sorting/sort_go	41.343s
```

Golang standard library implementation of [sort](https://cs.opensource.google/go/go/+/refs/tags/go1.17:src/sort/sort.go;drc=refs%2Ftags%2Fgo1.17;bpv=1;bpt=1;l=196?gsn=quickSort&gs=kythe%3A%2F%2Fgolang.org%3Flang%3Dgo%3Fpath%3Dsort%23func%2520quickSort) gives a lot of ideas about how to go about sorting choosing appropriate sorting algorithms for different size of input slices.
