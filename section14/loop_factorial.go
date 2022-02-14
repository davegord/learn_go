package main

import "fmt"

func main() {
	x := 4
	fac := loopFac(x)

	fmt.Println(fac)

}

func loopFac(x int) int {
	total := 1
	for ; x > 0; x-- {
		total *= x
	}
	return total
}
