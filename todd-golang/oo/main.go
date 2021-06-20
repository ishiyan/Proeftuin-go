package main

import (
	"fmt"
	"math"
)

// https://play.golang.org/p/jqKo448mg8w

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

func info(s shape) {
	fmt.Printf("%T, area is %v\n", s, s.area())
}

func main() {
	c := circle{
		radius: 10,
	}
	s := square{
		side: 10,
	}
	info(c)
	info(s)
}
