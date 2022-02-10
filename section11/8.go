package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:  []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`money_penny`: []string{`James Bond`, `Literature`, `Computer Science`},
	}

	x["dave_gordon"] = []string{"Jess Gordon", "Wine", "Dogs"}

	for i, v := range x {
		//fmt.Println(i, v)
		fmt.Println("This is the value for key", i)
		for i2, v2 := range v {
			fmt.Println("\t", i2, v2)
		}
	}

	delete(x, "dave_gordon")

	for i, v := range x {
		//fmt.Println(i, v)
		fmt.Println("This is the value for key", i)
		for i2, v2 := range v {
			fmt.Println("\t", i2, v2)
		}
	}

}
