# Naive implementations of sorting in Go vs. Standard Library

```
goos: darwin
goarch: arm64
pkg: algobook/ch2/sorting/sort_go
BenchmarkInsertionSort/Standard128-8         	  290389	      3964 ns/op	      24 B/op	       1 allocs/op
BenchmarkInsertionSort/Range128-8            	  232286	      5168 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort/Indices128-8          	  231967	      5171 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort/Standard1024-8        	   25586	     45738 ns/op	      24 B/op	       1 allocs/op
BenchmarkInsertionSort/Range1024-8           	    4914	    241729 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort/Indices1024-8         	    4940	    242073 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort/Standard65536-8       	     247	   4875175 ns/op	      24 B/op	       1 allocs/op
BenchmarkInsertionSort/Range65536-8          	       1	3123711166 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort/Indices65536-8        	       1	3119624500 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	algobook/ch2/sorting/sort_go	44.520s
```

Golang standard library implementation of [sort](https://cs.opensource.google/go/go/+/refs/tags/go1.17:src/sort/sort.go;drc=refs%2Ftags%2Fgo1.17;bpv=1;bpt=1;l=196?gsn=quickSort&gs=kythe%3A%2F%2Fgolang.org%3Flang%3Dgo%3Fpath%3Dsort%23func%2520quickSort) gives a lot of ideas about how to go about sorting choosing appropriate sorting algorithms for different size of input slices.
