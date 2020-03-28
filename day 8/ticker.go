package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 1600) //只打点3次
	ticker.Stop()
	fmt.Println("ticker stoped")
}
