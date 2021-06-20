package main

import (
	"fmt"
	"strconv"
)

// a struct is just a collect of fields
// in the following code, we define a struct called "Person", and then initialize it in two ways
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Go does not have classes. However, we can define methods on types.
// A method is a function with a special receiver argument.
// The following example shows how we use (or create) the method of the struct.
// In the code, we defined a method called (p Person) hello().
// A receiver in Go terms looks like this:
// func (receiver) func_name(parameters) return_type { code }

// method (value receiver)
func (p Person) hello() string {
	return "Hello, I am " + p.firstName + " " + p.lastName + ", " + strconv.Itoa(p.age) + " years old"
}

// The next example shows how we modified the data of the struct. In the code, we defined a method (receiver)
// called (p *Person) hasBirthday(). Note that it's taking a pointer.
// method (pointer receiver) -- modifies data
func (p *Person) hasBirthday() {
	p.age++
}

type Animal struct {
	Name string
	Age  int
}

// We made a method called String() which enables us to customize print for the Animal struct.
// String makes Animal satisfy the Stringer interface
func (a Animal) String() string {
	return fmt.Sprintf("%v %d", a.Name, a.Age)
}

type IPAddr [4]byte

func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

type Cities struct {
	name     string
	location [2]int
}

func main() {
	p1 := Person{firstName: "Steven", lastName: "King", city: "Chicago", gender: "m", age: 23}
	p2 := Person{"Neena", "Kochhar", "Boston", "f", 13}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.firstName, p2.lastName)
	// note that we can not only access a specific field of a struct but also we can modify the value of any field
	p2.age++
	fmt.Println(p2)

	// method (value receiver)
	fmt.Println(p1.hello())
	fmt.Println(p2.hello())

	// method (pointer receiver) -- modifies data
	p2.hasBirthday() // age + 1
	fmt.Println(p2.hello())
	p2.hasBirthday() // age + 1
	fmt.Println(p2.hello())

	a := Animal{Name: "Gopher", Age: 2}
	my_str := a.String()
	fmt.Println(my_str)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {1, 2, 3, 4},
	}
	for name, ip := range hosts {
		fmt.Printf("%v %v\n", name, ip)
	}

	// Create an empty slice of struct pointers
	cities := []*Cities{}
	// Create struct and append it to the slice
	ct := new(Cities)
	ct.name = "London"
	ct.location[0] = 5
	ct.location[1] = 0
	cities = append(cities, ct)
	// Create another struct
	ct = new(Cities)
	ct.name = "Sydney"
	ct.location[0] = 34
	ct.location[1] = 51
	cities = append(cities, ct)

	for i := range cities {
		c := cities[i]
		fmt.Println("City: ", *c)
	}
}
