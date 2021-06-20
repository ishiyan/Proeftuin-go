package main

import "fmt"

// Person is a person
type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old.\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I'm too tired to talk to you.\n")
}

// note no * in front of Person, because it is an interface
// note & in front of person, we return a pointer

// NewPerson creates a person
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p1 := NewPerson("James", 34)
	p1.SayHello()

	p2 := NewPerson("Jill", 134)
	p2.SayHello()
}
