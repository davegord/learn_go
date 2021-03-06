package main

import (
	"fmt"
	"sync"
)

var incrementor int

func main() {

	gs := 100

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			t := incrementor
			//fmt.Println("testing")
			t++
			incrementor = t
			wg.Done()
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("The value of incrementor when finsihed is", incrementor)

}
