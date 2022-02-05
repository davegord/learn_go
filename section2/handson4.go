package main

import "fmt"

type dave int

var x dave

func main() {
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	x = 42
	fmt.Println(x)
}
