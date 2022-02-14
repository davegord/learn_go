package main

import "fmt"

func main() {

	x := foo()

	y := x(3, 4)

	fmt.Printf("The type of x is %T\n", x)
	fmt.Println("The value of y is ", y)
	fmt.Printf("The type of y is %T\n", y)

}

func foo() func(x int, y int) int {
	return func(x1 int, y1 int) int {
		return x1 * y1
	}
}
