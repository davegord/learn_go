package main

import "fmt"

const (
	_ int = (iota + 2022)
	a
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
