package main

import "fmt"

func insertionSort(l []int) {
	var j int
	for i := range l {
		j = i
		for (j > 0) && (l[j] < l[j-1]) {
			l[j], l[j-1] = l[j-1], l[j]
			j = j - 1
		}
	}
}

func main() {
	fmt.Println("running")
	var l = []int{5, 4, 3, 2, 1}
	fmt.Println(l)
	insertionSort(l)
	fmt.Println(l)
}
