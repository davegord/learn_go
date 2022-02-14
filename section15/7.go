package main

import "fmt"

func main() {

	test := define

	fmt.Printf("The type of test is %T\n", test)
	fmt.Println(test(9, 10))

}

func define(i1 int, i2 int) int {
	return i1 + i2
}
