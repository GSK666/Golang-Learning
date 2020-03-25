package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sq := &s
	fmt.Println(sq.age)

	sq.age = 51
	fmt.Println(sq.age)
}
