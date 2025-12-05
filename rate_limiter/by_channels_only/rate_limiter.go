package ratelimiter

import (
	"fmt"
	"time"
)

func BufferedLimiter(n int, interval time.Duration) []time.Time {
	requests := make(chan int, 5)

	for i := range n {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(interval)
	var times []time.Time

	for req := range requests {
		<-limiter

		t := time.Now()
		times = append(times, t)
		fmt.Println("request ", req, t)
	}

	return times
}

func UnbufferedLimiter(n int, interval time.Duration) []time.Time {
	var times []time.Time
	requests := make(chan int)

	limiter := time.Tick(interval)

	go func() {
		for range n {
			requests <- n
		}
		close(requests) // if never closed, program will hang up, looking for more msgs
	}()

	for req := range requests {
		<-limiter

		t := time.Now()
		times = append(times, t)
		fmt.Println("request: ", req, t)
	}

	return times
}
