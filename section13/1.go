package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
	//fav string
}

func main() {

	p1 := person{
		first:      "Dave",
		last:       "Gordon",
		favFlavors: []string{"vanilla", "chocolate", "swirl"},
	}

	p2 := person{
		first:      "Jessica",
		last:       "Gordon",
		favFlavors: []string{"Chocolate", "Peppermind", "Coffee", "Raspberry"},
	}

	fmt.Printf("%s\t%s\t", p1.first, p1.last)

	for _, v := range p1.favFlavors {
		fmt.Printf("%s\t", v)
	}

	fmt.Printf("\n%s\t%s\t", p2.first, p2.last)
	for _, v := range p2.favFlavors {
		fmt.Printf("%s\t", v)
	}
	fmt.Println("")
}
