package main

import "fmt"

type person struct {
	first   string
	last    string
	friends []string
	phone   map[string]string
}

func main() {

	dave := person{
		first: "dave",
		last:  "gordon",
		friends: []string{
			"charles", "neil", "jon", "duane", "wally", "joe",
		},
		phone: map[string]string{
			"home": "919-460-5383",
			"cell": "919-280-4404",
			"work": "919-280-4404",
		},
	}

	fmt.Println(dave)

	for i, v := range dave.friends {
		fmt.Println(i, v)
	}

	for i, v := range dave.phone {
		fmt.Println(i, v)
	}
}
