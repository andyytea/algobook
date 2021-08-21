# Naive implementations of sorting in Go vs. Standard Library

```
goos: darwin
goarch: arm64
pkg: algobook/ch2/sorting/sort_go
BenchmarkSortIntStandard1K
BenchmarkSortIntStandard1K-8    	   25137	     45357 ns/op
BenchmarkInsertionSortN1K
BenchmarkInsertionSortN1K-8     	    4971	    242290 ns/op
BenchmarkSortIntStandard64K
BenchmarkSortIntStandard64K-8   	     235	   5080931 ns/op
BenchmarkInsertionSortN64K
BenchmarkInsertionSortN64K-8    	       2	 610415958 ns/op
PASS
ok  	algobook/ch2/sorting/sort_go	7.948s
```
