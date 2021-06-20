package main

import "fmt"

type employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	developer = iota
	manager
)

// functional
func newEmployee(role int) *employee {
	switch role {
	case developer:
		return &employee{"", "Developer", 60000}
	case manager:
		return &employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := newEmployee(manager)
	m.Name = "Sam"
	fmt.Println(m)
}
