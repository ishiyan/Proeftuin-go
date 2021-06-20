package main

import "fmt"

// Person represents a person.
type Person struct {
	name, position string
}

type personMod func(*Person)

// PersonBuilder is a builder.
type PersonBuilder struct {
	actions []personMod
}

// Called sets in a person's name.
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

// Build builds an object.
func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

// extend PersonBuilder

// WorksAsA sets a person's position.
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Dmitri").WorksAsA("dev").Build()
	fmt.Println(*p)
}
