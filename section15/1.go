package main

import (
	"fmt"
)

func main() {
	a := foo()
	fmt.Println("Outisde foo a is", a)

	b, c := bar()
	fmt.Println("Outside bar b and c are", b, c)

}

func foo() int {
	fmt.Println("Inside foo")
	x := 2000
	return x
}

func bar() (int, string) {
	st := "Outside bar"
	x := 42
	return x, st

}
