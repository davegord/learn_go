package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var incrementor int

func main() {

	gs := 100

	var wg sync.WaitGroup

	var incrementor int64

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			//fmt.Println("testing")
			atomic.AddInt64(&incrementor, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("The value of incrementor when finsihed is", incrementor)

}
