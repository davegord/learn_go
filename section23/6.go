package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)
	receive(c)

}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Done!!")

}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)

}
