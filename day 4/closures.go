package main

import "fmt"

func intseq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextint := intseq()

	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())

	newint := intseq()
	fmt.Println(newint())
}
