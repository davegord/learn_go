package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := foo(s...)
	fmt.Println(sum)

	sum_slice := bar(s)
	fmt.Println(sum_slice)

}

func foo(x ...int) int {
	var ps int

	for _, v := range x {
		ps += v
	}
	return ps
}

func bar(si []int) int {
	var sum int

	for _, v := range si {
		sum += v
	}
	return sum
}
