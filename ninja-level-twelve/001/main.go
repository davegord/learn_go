package main

import (
	"fmt"

	"github.com/davegord/learn_go/ninja-level-twelve/001/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(7),
	}

	fmt.Println(fido)

}
