package main

import "fmt"

func main() {
	x := 10
	y := 20

	fmt.Println("before one")
	defer one(x, y)
	two()

}

func one(x int, y int) {
	fmt.Println("I'm in one")
	fmt.Println(x + y)
}

func two() {
	fmt.Println("I'm in two, I was called after run, but one was deferred so I'm running first")
}
