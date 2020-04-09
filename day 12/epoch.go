package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sec := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(sec)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(sec, 0))
	fmt.Println(time.Unix(0, nanos))
}
