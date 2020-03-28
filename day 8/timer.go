package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second * 2)

	<-time1.C
	fmt.Println("timer1 expired")

	time2 := time.NewTimer(time.Second) //还没失效就结束了
	go func() {
		<-time2.C
		fmt.Println("timer2 expired")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("time2 stoped")
	}
}
