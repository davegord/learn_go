package main

import "fmt"

type dave int

var y int

var x dave

func main() {
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("The type of y is %T\n", y)

}
