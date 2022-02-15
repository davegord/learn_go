package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	//	y := foo(bar(x))

	//	fmt.Println(y)
}

/*
func foo(f func(xi []int) int) int {
	n := f()
	n++
	return n
}

func bar(xs []int) int {

	x := 0

	for _, v := range xs {
		x += v
	}
	return x
}
*/
