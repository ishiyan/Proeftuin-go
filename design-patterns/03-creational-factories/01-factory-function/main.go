package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{name, age}
}

func main() {
	// initialize directly
	p := person{"John", 22}
	fmt.Println(p)

	// use a constructor
	p2 := newPerson("Jane", 21)
	p2.age = 30
	fmt.Println(p2)
}
