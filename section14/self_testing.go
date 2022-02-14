package main

import "fmt"

func main() {

	s := test()
	fmt.Println("This is what s equals", s)

	fmt.Println("----------------")

	sum_func := test_param(20, 30)

	fmt.Println("This is the sum of the function", sum_func)

}

func test() string {
	s := "Hello world"
	return s
}

func test_param(i1 int, i2 int) int {
	sum := i1 + i2
	return sum
}
