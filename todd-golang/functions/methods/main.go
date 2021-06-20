package main

import (
	"fmt"
)

// https://play.golang.org/p/uVDLbDiLR4i

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	p.speak()
	p2.speak()
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, p.age, "years old")
}
