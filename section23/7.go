package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
		//fmt.Println(i)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c, k)
	}

}

/*
func send(c chan<- int) {
	//fmt.Println("In send")
	for i := 0; i < 10; i++ {
		c <- i
	}
}
*/
