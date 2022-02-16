package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first)
	fmt.Println(p.last)
	fmt.Println(p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	d := person{
		first: "Dave",
		last:  "Gordon",
		age:   42,
	}

	//d.speak()

	fmt.Println("Trying to call saySomething with a pointer")
	saySomething(&d)

	fmt.Println("Trying to call something without a pointer")
	saySomething(d)
}
