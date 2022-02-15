package main

import "fmt"

type person struct {
	first   string
	last    string
	address string
}

func changeMe(p *person) {
	(*p).address = "103 cumberland green drive"
}

func main() {
	dave := person{
		first:   "dave",
		last:    "gordon",
		address: "104 bending oak way",
	}

	fmt.Println(dave)

	changeMe(&dave)

	fmt.Println(dave)
}
