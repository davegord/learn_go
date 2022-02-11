package main

import "fmt"

func main() {
	x := struct {
		first       string
		friends     map[string]int
		instruments []string
	}{
		first: "dave",
		friends: map[string]int{
			"joe":     1,
			"charles": 2,
			"neil":    3,
			"wally":   4,
		},
		instruments: []string{
			"gibson les paul", "hofner bass", "grand piano",
		},
	}

	fmt.Println(x)
}
