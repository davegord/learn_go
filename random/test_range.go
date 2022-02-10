package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println("The value of index is", i, "The value of value is", v)
	}

	//slice a slice!

	y := x[2:4]

	fmt.Println("This is y", y)

	//append to a slice

	x = append(x, 80)
}
