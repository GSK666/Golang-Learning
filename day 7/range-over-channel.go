package main

import "fmt"

func main() {
	quene := make(chan string, 2)
	quene <- "one"
	quene <- "two"
	close(quene)

	for elem := range quene {
		fmt.Println(elem)
	}
}
