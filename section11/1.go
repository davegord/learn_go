package main

import "fmt"

func main() {
	x := [5]int{42, 43, 44, 45, 46}

	//fmt.Println(x)
	for i, v := range x {
		fmt.Printf("this is index %v, %v\n", i, v)
	}

	fmt.Printf("%T\n", x)
}
