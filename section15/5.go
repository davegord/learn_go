package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	area := s.length * s.length
	return area
}

func (c circle) area() float64 {
	pi := math.Pi
	area := pi * (c.radius * c.radius)
	return area
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		length: 15,
	}

	//a1 := s1.area()
	info(s1)

	//fmt.Println(a1)
	//info(s1)
	c1 := circle{
		radius: 12.345,
	}

	info(c1)

}
