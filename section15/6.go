package main

import "fmt"

func main() {

	x := func() int {
		return 9 * 9
	}()

	fmt.Println(x)
}
