package main

import "fmt"

// Employee is an employee
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// what if we want factories for specific roles?
// functional approach

// NewEmployeeFactory creates an employee factory
func NewEmployeeFactory(position string,
	annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach

// EmployeeFactory is an employee factory
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

// NewEmployeeFactory2 is an employee factory
func NewEmployeeFactory2(position string,
	annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

// Create creates an employee
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	fmt.Println(developer)

	manager := managerFactory("Jane")
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	// can modify post-hoc
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
