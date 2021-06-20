package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

type relationship int

const (
	aparent relationship = iota
	achild
	asibling
)

type person struct {
	name string
	// other useful stuff here
}

type info struct {
	from         *person
	relationship relationship
	to           *person
}

type relationshipBrowser interface {
	FindAllChildrenOf(name string) []*person
}

type relationships struct {
	relations []info
}

func (rs *relationships) FindAllChildrenOf(name string) []*person {
	result := make([]*person, 0)

	for i, v := range rs.relations {
		if v.relationship == aparent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}

	return result
}

func (rs *relationships) AddParentAndChild(parent, child *person) {
	rs.relations = append(rs.relations,
		info{parent, aparent, child})
	rs.relations = append(rs.relations,
		info{child, achild, parent})
}

type research struct {
	// relationships relationships
	browser relationshipBrowser // low-level
}

func (r *research) Investigate() {
	//relations := r.relationships.relations
	//for _, rel := range relations {
	//	if rel.from.name == "John" &&
	//		rel.relationship == Parent {
	//		fmt.Println("John has a child called", rel.to.name)
	//	}
	//}

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := person{"John"}
	child1 := person{"Chris"}
	child2 := person{"Matt"}

	// low-level module
	relationships := relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := research{&relationships}
	research.Investigate()
}
