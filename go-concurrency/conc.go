package main

import (
	"fmt"
	"sync"
	"time"
)

func pSeries(s int, num int, start time.Time, wg *sync.WaitGroup) {
	t := time.Now()
	fmt.Printf("series: %v \t number: %v \t time: %v\n", s, num, t.Sub(start))
	wg.Done()
}

func bigFuncConcurrent(s int) time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	for s > 0 {
		for i := 0; i <= 1000; i++ {
			wg.Add(1)
			go pSeries(s, i, start, &wg)
		}
		s--
	}
	wg.Wait()
	end := time.Now()
	return end.Sub(start)
}

func main() {
	const ITERATIONS = 50
	const NSERIES = 5
	tarr := make([]time.Duration, ITERATIONS)
	var sum time.Duration
	for i := 0; i < ITERATIONS; i++ {
		t := bigFuncConcurrent(NSERIES)
		tarr[i] = t
		sum += t
	}

	fmt.Println(tarr)
	fmt.Printf("Average Time: %v\n", sum/ITERATIONS)
}
