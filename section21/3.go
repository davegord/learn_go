package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementor int

func main() {

	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			t := incrementor
			//fmt.Println("testing")
			runtime.Gosched()
			t++
			incrementor = t
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("The value of incrementor when finsihed is", incrementor)

}
