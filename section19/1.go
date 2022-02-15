package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	Users := []User{u1, u2, u3}

	fmt.Println(Users)

	// your code goes here
	s, err := json.Marshal(Users)

	if err != nil {
		fmt.Println("There was an error")
	}

	fmt.Println(string(s))
}
