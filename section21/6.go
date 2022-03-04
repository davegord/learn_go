package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This arch is", runtime.GOARCH)
	fmt.Println("This os is", runtime.GOOS)

}
