package main

import (
	"fmt"
	"time"
)

func pSeriesSync(s int, num int, start time.Time) {
	t := time.Now()
	fmt.Printf("series: %v \t number: %v \t time: %v\n", s, num, t.Sub(start))
}

func bigFuncSynchronous(s int) time.Duration {
	start := time.Now()
	for s > 0 {
		for i := 0; i <= 10000; i++ {
			pSeriesSync(s, i, start)
		}
		s--
	}
	end := time.Now()
	return end.Sub(start)
}

func main() {

	const ITERATIONS = 50
	const NSERIES = 5
	tarr := make([]time.Duration, ITERATIONS)
	var sum time.Duration

	for i := 0; i < ITERATIONS; i++ {
		t := bigFuncSynchronous(NSERIES)
		tarr[i] = t
		sum += t
	}

	fmt.Println(tarr)
	fmt.Printf("Average Time: %v\n", sum/ITERATIONS)
}
