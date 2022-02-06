package main

import "fmt"

func main() {
	dave := "dave"

	if dave == "dave" {
		fmt.Println("equals dave")
	} else if dave == "Joe" {
		fmt.Println("equals Joe")
	} else {
		fmt.Println("There was no match")
	}
}
