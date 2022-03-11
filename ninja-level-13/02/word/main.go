package word

import (
	"fmt"
	"strings"
)

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	//fmt.Println(xs)

	/*
		counter := 0
		for i, v := range xs {
			counter += 1
			fmt.Println(i, v)
		}

	*/
	counter := len(xs)

	fmt.Println("---------")

	return counter
}
