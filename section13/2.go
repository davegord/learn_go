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
		last:       "Walpole Gordon",
		favFlavors: []string{"Chocolate", "Peppermind", "Coffee", "Raspberry"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for j, k := range v.favFlavors {
			fmt.Println(j, k)
		}

	}

	/*
		fmt.Printf("%s\t%s\t", p1.first, p1.last)

		for _, v := range p1.favFlavors {
			fmt.Printf("%s\t", v)
		}

		fmt.Printf("\n%s\t%s\t", p2.first, p2.last)
		for _, v := range p2.favFlavors {
			fmt.Printf("%s\t", v)
		}
		fmt.Println("")
	*/
}
