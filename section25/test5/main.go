package main

import "fmt"

func main() {
	fmt.Println("Here we go")
	Foo(5, 10)

}

func Foo(i, j int) int {
	s := i + j
	//fmt.Printf("In Foo, the sum is %v\n", s)
	return s
}
