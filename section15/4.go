package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (pe person) speak() {
	fmt.Println("My name is", pe.first, pe.last, "and my age is", pe.age)
}

func main() {

	dave := person{
		first: "Dave",
		last:  "Gordon",
		age:   42,
	}

	dave.speak()
}
