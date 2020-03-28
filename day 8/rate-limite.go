package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstylimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstylimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstylimiter <- t
		}
	}()

	burstylimiterquests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstylimiterquests <- i
	}
	close(burstylimiterquests)
	for req := range burstylimiterquests {
		<-burstylimiterquests
		fmt.Println("request", req, time.Now())
	}
}
