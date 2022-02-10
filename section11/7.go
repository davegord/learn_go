package main

import "fmt"

func main() {

	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	ss := [][]string{s1, s2}

	//fmt.Println(s[1])

	for i, v := range ss {
		fmt.Println("for record: ", i)
		for j, w := range v {
			fmt.Println(j, w)
		}

	}
}
