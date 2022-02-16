package main

import (
	"fmt"
	"runtime"
	"sync"
)

var w sync.WaitGroup

func main() {

	fmt.Println("Runtime\t\t", runtime.GOOS)
	fmt.Println("Arch\t\t", runtime.GOARCH)
	fmt.Println("Numcpu\t\t", runtime.NumCPU())
	fmt.Println("goroutines\t", runtime.NumGoroutine())

	w.Add(2)

	go foo()
	go bar()

	fmt.Println("goroutines\t", runtime.NumGoroutine())
	w.Wait()
}

func foo() {
	fmt.Println("In foo")
	w.Done()
}

func bar() {
	fmt.Println("In bar")
	w.Done()
}
