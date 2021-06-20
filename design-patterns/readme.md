# Design Patterns in Go

Discover the modern implementation of design patterns in Go (golang).

## Introduction

- Design patterns are typically OOP-based.
- Go is not OOP
  - no inheritance
  - weak incapsulation
- Permissive visibility/naming
- Use of OOP-ish terminology
  - hierarchy implies a set of related types: shared interface, embedding
  - properties are combinations of getter/setter methods

## SOLID Design Principles

Useful principles of object-oriented design, [wikipedia.org/wiki/SOLID](https://en.wikipedia.org/wiki/SOLID).
Introduced by Robert C. Martin (uncle Bob).

- **S**ingle Responsibility Principle (SRP)
- **O**pen-Closed Principle (OCP)
- **L**iskov Substitution Principle (LSP)
- **I**nterface Segregation Principle (ISP)
- **D**ependency Inversion Principle (DIP)

### SRP - Single Responsibility Principle

Each type should have primary responsibility and as a result one reason to change.
Segregation of Concerns.
Antipattern: God Object, you put everything into it.

```go
type Journal struct {
  entries []string
}
func (j *Journal) String() string {
  // ...
}
func (j *Journal) AddEntry(text string) int {
  // ...
}
func (j *Journal) RemoveEntry(index int) {
  // ...
}

// breaks srp

func (j *Journal) Save(filename string) {
  // ...
}
func (j *Journal) Load(filename string) {
  // ...
}
func (j *Journal) LoadFromWeb(url *url.URL) {
  // ...
}

// fix it by making a simple function

func SaveToFile(j *Journal, filename string) {
  // ...
}

// fif it by putting into different package

type Persistence struct {
  lineSeparator string
}
func (p *Persistence) saveToFile(j *Journal, filename string) {
  // ...
}
```

### OCP - Open-Closed Principle

Open for extension, closed for modification.
Also Enterprise pattern called `Specification` which illustrates OCP very well.

```go
type Color int
const (
  red Color = iota
  green
  blue
)

type Size int
const (
  small Size = iota
  medium
  large
)

type Product struct {
  name string
  color Color
  size Size
}
type Filter struct {
}

// this violates OCP, since you have to modify for every new filtering type

func (f *Filter) filterByColor(products []Product, color Color)[]*Product {
}
func (f *Filter) filterBySize(products []Product, size Size) []*Product {
}
func (f *Filter) filterBySizeAndColor(products []Product, size Size, color Color)[]*Product {
}

// correct implementation

type Specification interface {
  IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
  color Color
}
func (spec ColorSpecification) IsSatisfied(p *Product) bool {
  return p.color == spec.color
}

type SizeSpecification struct {
  size Size
}
func (spec SizeSpecification) IsSatisfied(p *Product) bool {
  return p.size == spec.size
}

type AndSpecification struct {
  first, second Specification
}
func (spec AndSpecification) IsSatisfied(p *Product) bool {
  return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type BetterFilter struct {}
func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
}

greenSpec := ColorSpecification{green}
largeSpec := SizeSpecification{large}
largeGreenSpec := AndSpecification{largeSpec, greenSpec}
bf := BetterFilter{}
bf.Filter(products, largeGreenSpec)
```

### LSP - Liskov Substitution Principle

[Barbara Liskov](https://en.wikipedia.org/wiki/Barbara_Liskov)
If you have functionality which works correctly with a base class, it should also work correctly with a dervide class.
Since we don't have classes in Go we''l explore a variation of this.
Implementations of generalizations should not break functionality at the base level.

**Modified LSP**
*If a function takes an interface and works with a type T that implements this interface, any structure that aggregates T should also be usable in that function.*

```go
type Sized interface {
  GetWidth() int
  SetWidth(width int)
  GetHeight() int
  SetHeight(height int)
}

type Rectangle struct {
  width, height int
}
func (r *Rectangle) GetWidth() int {
  return r.width
}
func (r *Rectangle) SetWidth(width int) {
  r.width = width
}
func (r *Rectangle) GetHeight() int {
  return r.height
}
func (r *Rectangle) SetHeight(height int) {
  r.height = height
}

// aggregates Rectangle
type Square struct {
  Rectangle
}
// violates LSP
func (s *Square) SetWidth(width int) {
  s.width = width
  s.height = width
}
// violates LSP
func (s *Square) SetHeight(height int) {
  s.width = height
  s.height = height
}

// one of ways to solve it: still use Rectangle functionality
type Square2 struct {
  size int
}
func (s *Square2) Rectangle() Rectangle {
  return Rectangle{s.size, s.size}
}
```

### ISP - Interface Segregation Principle

The simplest principle. What it says you should not put too much in an interface.
It is better to beak it into smaller interfaces.

```go
type Document struct {}

type Machine interface {
  Print(d Document)
  Fax(d Document)
  Scan(d Document)
}

// ok if you need a multifunction device

type MultiFunctionPrinter struct {}
func (m MultiFunctionPrinter) Print(d Document) {}
func (m MultiFunctionPrinter) Fax(d Document) {}
func (m MultiFunctionPrinter) Scan(d Document) {}

// here you have to define methods you cannot support

type OldFashionedPrinter struct {}
func (o OldFashionedPrinter) Print(d Document) {}
func (o OldFashionedPrinter) Fax(d Document) {
  panic("operation not supported")
}
func (o OldFashionedPrinter) Scan(d Document) {
  panic("operation not supported")
}

// better approach: split into several interfaces

type Printer interface {
  Print(d Document)
}
type Fax interface {
  Fax(d Document)
}
type Scanner interface {
  Scan(d Document)
}

// printer only

type MyPrinter struct {}
func (m MyPrinter) Print(d Document) {}

// combine interfaces

type Photocopier struct {}
func (p Photocopier) Scan(d Document) {}
func (p Photocopier) Print(d Document) {}

type MultiFunctionDevice interface {
  Printer
  Scanner
  Fax
}

// interface combination + decorator

type MultiFunctionMachine struct {
  printer Printer
  fax     Fax
  scanner Scanner
}
func (m MultiFunctionMachine) Print(d Document) {
  m.printer.Print(d)
}
func (m MultiFunctionMachine) Fax(d Document) {
  m.fax.Fax(d)
}
func (m MultiFunctionMachine) Scan(d Document) {
  m.scanner.Scan(d)
}
```

### DIP - Dependency Inversion Principle

**Not** Dependency Injection, however they have relationship.
Dependency Injection is one of possible **implementations** of DIP.

High Level Modules should not depend on Low Level Modules.
Bith should depend on abstractions.

Let's do a genealogy research and explore relationships between people.

```go
type Relationship int
const (
  Parent Relationship = iota
  Child
  Sibling
)

type Person struct {
  name string
}

type Info struct {
  from *Person
  relationship Relationship
  to *Person
}
// Low Level Module
type Relationships struct {
  relations []Info
}
func (rs *Relationships) AddParentAndChild(parent, child *Person) {
  // ...
}

// High Level Module
type Research struct {
  relationships Relationships // low level dependency, violates DIP
}
func (r *Research) Investigate() {
  relations := r.relationships.relations
  // ...
}

// Both should depend on abstraction

type RelationshipBrowser interface {
  FindAllChildrenOf(name string) []*Person
}
// Low Level Module implements interface
func (rs *Relationships) FindAllChildrenOf(name string) []*Person {}

// High Level Module depends only on interface
type Research2 struct {
  browser RelationshipBrowser
}
func (r *Research2) Investigate() {
  for _, p := range r.browser.FindAllChildrenOf("John") {
    fmt.Println("John has a child called", p.name)
  }
}
```

## Creational Design Patterns

### Builder -- when construction gets a little bit too complicated

#### Builder

Motivation

- Some objects are simple and can be createdin a single construction call.
- Other objects require a bit of ceremony to create.
- Having a factory function with 10 arguments is not productive, you force a customer to make 10 decisions in a single call.
- Instead, opt for a piecewise (piece-by-piece) construction.
- Builder provides an API for contructing an object step-by-step.

When piecewise construction is complicated, provide an API for doing it succinctly.

Let's look at a builder which is already built into Go -- `strings.Builder`.
Let's pretend building an HTML web server which builds HTML from the text pieces.

```go
  hello := "hello"
  sb := strings.Builder{}
  sb.WriteString("<p>")
  sb.WriteString(hello)
  sb.WriteString("</p>")
  fmt.Printf("%s\n", sb.String())

  // suppose we have several words and we want to put them into HTML list
  // observe we have 3 distinct parts here
  words := []string{"hello", "world"}
  sb.Reset()
  // '<ul><li>...</li><li>...</li><li>...</li></ul>'
  sb.WriteString("<ul>") // PART 1
  for _, v := range words { // PART 2
    sb.WriteString("<li>")
    sb.WriteString(v)
    sb.WriteString("</li>")
  }
  sb.WriteString("</ul>") // PART 3
  fmt.Println(sb.String())
```

Let's create a builder.

```go
type HtmlElement struct {
  name, text string
  elements   []HtmlElement
}
func (e *HtmlElement) String() string {}

type HtmlBuilder struct {
  rootName string
  root     HtmlElement
}
func NewHtmlBuilder(rootName string) *HtmlBuilder {
  b := HtmlElement{rootName, "", []HtmlElement{}}}
  return &b
}
func (b *HtmlBuilder) String() string {
  return b.root.String()
}
func (b *HtmlBuilder) AddChild(childName, childText string) {
  e := HtmlElement{childName, childText, []HtmlElement{}}
  b.root.elements = append(b.root.elements, e)
}

// usage
b := NewHtmlBuilder("ul")
b.AddChild("li", "hello")
b.AddChild("li", "world")
fmt.Println(b.String())
```

Fluent Interface:

- Allows to chain calls.
- Returns a pointer to receiver at the end of the method.

```go
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
  e := HtmlElement{childName, childText, []HtmlElement{}}
  b.root.elements = append(b.root.elements, e)
  return b
}

// usage
b := NewHtmlBuilder("ul")
b.AddChildFluent("li", "hello").AddChildFluent("li", "world")
fmt.Println(b.String())
```

#### Builder Facets

Normally a single builder is sufficient to build a particular object.
But there are situations when you need more that one builder.
For example, when you need to separate different aspects for a Builder of a particulat type.

Let's look at example of two aspects of a person -- address and job.
We can create a basic `PersonBuilder` which will create an empty `Person`.
Then we make additional builders which will aggregate the `PersonBuilder`.
We provide two utility methods in the `PersonBuilder` which return an `AddressBuilder` and a `JobBuilder`.

Effectively, we create a tiny DSL for building a person.
All 3 builders work in cooperation because they share the same pointer to a person.

```go
type Person struct {
  StreetAddress, Postcode, City string
  CompanyName, Position string
  AnnualIncome int
}

type PersonBuilder struct {
  person *Person // needs to be inited
}
func NewPersonBuilder() *PersonBuilder {
  return &PersonBuilder{&Person{}}
}
func (it *PersonBuilder) Build() *Person {
  return it.person
}
func (it *PersonBuilder) Works() *PersonJobBuilder {
  return &PersonJobBuilder{*it}
}
func (it *PersonBuilder) Lives() *PersonAddressBuilder {
  return &PersonAddressBuilder{*it}
}

type PersonJobBuilder struct {
  PersonBuilder
}
func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
  pjb.person.CompanyName = companyName
  return pjb
}
func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
  pjb.person.Position = position
  return pjb
}
func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
  pjb.person.AnnualIncome = annualIncome
  return pjb
}

type PersonAddressBuilder struct {
  PersonBuilder
}
func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
  pab.person.StreetAddress = streetAddress
  return pab
}
func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
  pab.person.City = city
  return pab
}
func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
  pab.person.Postcode = postcode
  return pab
}

// usage
pb := NewPersonBuilder()
pb.
  Lives().
    At("123 London Road").
    In("London").
    WithPostcode("SW12BC").
  Works().
    At("Fabrikam").
    AsA("Programmer").
    Earning(123000)
person := pb.Build()
fmt.Println(*person)
```

#### Builder Parameter

Hide your objects inside the builder and provide utility metods to manipulate them.

Builder Parameter is a function which applies to the builder.

```go
type email struct {
  from, to, subject, body string
}

type EmailBuilder struct {
  email email
}
func (b *EmailBuilder) From(from string) *EmailBuilder {
  b.email.from = from
  return b
}
func (b *EmailBuilder) To(to string) *EmailBuilder {
  b.email.to = to
  return b
}
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
  b.email.subject = subject
  return b
}
func (b *EmailBuilder) Body(body string) *EmailBuilder {
  b.email.body = body
  return b
}

func sendMailImpl(email *email) { // private implementation
  // actually ends the email
}

type build func(*EmailBuilder)

// action is a builder parameter
func SendEmail(action build) {
  builder := EmailBuilder{}
  action(&builder)
  sendMailImpl(&builder.email)
}

// usage
SendEmail(func(b *EmailBuilder) {
  b.From("foo@bar.com").
    To("bar@baz.com").
    Subject("Meeting").
    Body("Hello, do you want to meet?")
)
```

#### Functional Builder

You can extend already existing builder by using a functional programming approach.

Imagine we have a builder for an aspect (person's name) of a type person and we want to extend it for a person's position.
Builder contains a list of actions (modifications) which apply to the default person during the Build call.
It represents a *delayed* builder.

```go
type Person struct {
  name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
  actions []personMod
}
func (b *PersonBuilder) Called(name string) *PersonBuilder {
  b.actions = append(b.actions, func(p *Person) {
    p.name = name
  })
  return b
}
func (b *PersonBuilder) Build() *Person {
  p := Person{}
  for _, a := range b.actions {
    a(&p)
  }
  return &p
}

// extend PersonBuilder
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
  b.actions = append(b.actions, func(p *Person) {
    p.position = position
  })
  return b
}

// usage
b := PersonBuilder{}
p := b.Called("James").WorksAsA("agent").Build()
fmt.Println(*p)
```

### Factories -- ways of controlling how an object is constructed

- Object creation logic becomes sometimes convoluted
- Struct may have too many fields, need to initialize all correctly
- Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to:
  - a separate function (Factory Function, a.k.a. Constructor), which is common approach in Go
  - that may exist in a separate struct (Factory)

Factory is a component responsible solely for the wholesale (in a single call, not piecewise) creation of objects.

#### Factory Function

Sometimes you need some behavior to happen when a struct is being created. For instance, you want to have a default value or to do an argument validation.

A `factory function` is a free-standing function which returns an instance of a struct.

It is slightly more efficient to return a pointer to the struct that to return it by value.

```go
package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func NewPerson(name string, age int) *Person {
  if age < 0 {
    panic()
  }
  return &Person{name, age}
}

func main() {
  // use a constructor
  p2 := NewPerson("Jane", 21)
  p2.Age = 30
  fmt.Println(p2)
}
```

#### Interface Factory

When you have a factory function, you don't always have to return a struct.
Instead, you can return an interface to which this struct conforms to.
Now you cannot modify an underlying type.
A neat way of encapsulation.

```go
package main

import "fmt"

type Person interface {
  SayHello()
}

type person struct {
  name string
  age int
}

type tiredPerson struct {
  name string
  age int
}

func (p *person) SayHello() {
  fmt.Printf("Hi, my name is %s, I am %d years old.\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
  fmt.Printf("Sorry, I'm too tired to talk to you.\n")
}

// note no * in front of Person, because it is an interface
// note & in front of person, we return a pointer
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
```

#### Factory Generator

Generates a factory. We may want to create factories dependent on parameters with ability to modify them.

There are two ways to do it.

- Functional. Return a function which creates a factory. An advantage is that you can pass this factory function into another functions as an argument, so users can just call it.
- Structural. Make a factory a struct having some kind of a `Create` method. Users can change properties since the factory is a struct, but they should know they have to call a `Create` method.

Functional way is more idiomatic.

```go
package main

import "fmt"

type Employee struct {
  Name, Position string
  AnnualIncome int
}

// functional approach
func NewEmployeeFactory(position string,
  annualIncome int) func(name string) *Employee {
  return func(name string) *Employee {
    return &Employee{name, position, annualIncome}
  }
}

// structural approach
type EmployeeFactory struct {
  Position string
  AnnualIncome int
}

func NewEmployeeFactory2(position string,
  annualIncome int) *EmployeeFactory {
  return &EmployeeFactory{position, annualIncome}
}

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
```

#### Prototype Factory

A factory function which operates on pre-configured objects.
Create constants identifying pre-defined objects, and pass these constants into a factory function.

```go
package main

import "fmt"

type Employee struct {
  Name, Position string
  AnnualIncome int
}

const (
  Developer = iota
  Manager
)

// functional
func NewEmployee(role int) *Employee {
  switch role {
  case Developer:
    return &Employee{"", "Developer", 60000}
  case Manager:
    return &Employee{"", "Manager", 80000}
  default:
    panic("unsupported role")
  }
}

func main() {
  m := NewEmployee(Manager)
  m.Name = "Sam"
  fmt.Println(m)
}
```

### Prototype -- when it's easier to copy an existing object to fully initialize a new one

- Complicated objects aren't design from scratch. They reiterate existing design.
- An existing (partially or fully constructed) design is a Prototype.
- We make a copy of the prototype and customize it. This requires *deep copy* support.
- We make the cloning convenient (e.g., via a Factory).

Prototype is a partially or fully initialized object that you copy (clone) and make use of.

- To implement a prototype, partially construct an object and store it somewhere.
- Deep copy the prototype.
- Customize the resulting instance.
- A prototype factory provides a convenient API for using prototypes.

#### Deep copying

```go
package prototype

import "fmt"

type Address struct {
  StreetAddress, City, Country string
}

type Person struct {
  Name string
  Address *Address
}

func main() {
  john := Person{"John", &Address{"123 London Rd", "London", "UK"}}

  //jane := john

  // shallow copy
  //jane.Name = "Jane" // ok

  //jane.Address.StreetAddress = "321 Baker St"

  //fmt.Println(john.Name, john.Address)
  //fmt.Println(jane.Name, jane. Address)

  // what you really want
  jane := john
  jane.Address = &Address{
    john.Address.StreetAddress,
    john.Address.City,
    john.Address.Country  }

  jane.Name = "Jane" // ok

  jane.Address.StreetAddress = "321 Baker St"

  fmt.Println(john.Name, john.Address)
  fmt.Println(jane.Name, jane. Address)
}
```

#### Copy Method

```go
package prototype

import "fmt"

type Address struct {
  StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
  return &Address{
    a.StreetAddress,
    a.City,
    a.Country }
}

type Person struct {
  Name string
  Address *Address
  Friends []string
}

func (p *Person) DeepCopy() *Person {
  q := *p // copies Name
  q.Address = p.Address.DeepCopy()
  copy(q.Friends, p.Friends)
  return &q
}

func main() {
  john := Person{"John",
    &Address{"123 London Rd", "London", "UK"},
    []string{"Chris", "Matt"}}

  jane := john.DeepCopy()
  jane.Name = "Jane"
  jane.Address.StreetAddress = "321 Baker St"
  jane.Friends = append(jane.Friends, "Angela")

  fmt.Println(john, john.Address)
  fmt.Println(jane, jane.Address)
}
```

#### Copy Through Serialization

```go
package main

import (
  "bytes"
  "encoding/gob"
  "fmt"
)

type Address struct {
  StreetAddress, City, Country string
}

type Person struct {
  Name string
  Address *Address
  Friends []string
}

func (p *Person) DeepCopy() *Person {
  // note: no error handling below
  b := bytes.Buffer{}
  e := gob.NewEncoder(&b)
  _ = e.Encode(p)

  // peek into structure
  fmt.Println(string(b.Bytes()))

  d := gob.NewDecoder(&b)
  result := Person{}
  _ = d.Decode(&result)
  return &result
}

func main() {
  john := Person{"John",
    &Address{"123 London Rd", "London", "UK"},
    []string{"Chris", "Matt", "Sam"}}

  jane := john.DeepCopy()
  jane.Name = "Jane"
  jane.Address.StreetAddress = "321 Baker St"
  jane.Friends = append(jane.Friends, "Jill")

  fmt.Println(john, john.Address)
  fmt.Println(jane, jane.Address)
}
```

#### Prototype Factory (2)

```go
package main

import (
  "bytes"
  "encoding/gob"
  "fmt"
)

type Address struct {
  Suite int
  StreetAddress, City string
}

type Employee struct {
  Name string
  Office Address
}

func (p *Employee) DeepCopy() *Employee {
  // note: no error handling below
  b := bytes.Buffer{}
  e := gob.NewEncoder(&b)
  _ = e.Encode(p)

  // peek into structure
  //fmt.Println(string(b.Bytes()))

  d := gob.NewDecoder(&b)
  result := Employee{}
  _ = d.Decode(&result)
  return &result
}

// employee factory
// either a struct or some functions
var mainOffice = Employee {
  "", Address{0, "123 East Dr", "London"}}
var auxOffice = Employee {
  "", Address{0, "66 West Dr", "London"}}

// utility method for configuring emp
//   ↓ lowercase
func newEmployee(proto *Employee,
  name string, suite int) *Employee {
  result := proto.DeepCopy()
  result.Name = name
  result.Office.Suite = suite
  return result
}

func NewMainOfficeEmployee(
  name string, suite int) *Employee {
    return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(
  name string, suite int) *Employee {
    return newEmployee(&auxOffice, name, suite)
}

func main() {
  // most people work in one of two offices

  //john := Employee{"John",
  //  Address{100, "123 East Dr", "London"}}
  //
  //jane := john.DeepCopy()
  //jane.Name = "Jane"
  //jane.Office.Suite = 200
  //jane.Office.StreetAddress = "66 West Dr"

  john := NewMainOfficeEmployee("John", 100)
  jane := NewAuxOfficeEmployee("Jane", 200)

  fmt.Println(john)
  fmt.Println(jane)
}
```

### Singleton -- a design pattern everyone loves to hate... but is it really that bad

A component which is instantiated only once.

When discussing which patterns to drop, we found that we still love them. (Not really -- I'm in favor of dropping Singleton. Its use is almost always a design smell.) *Erich Gamma*

Motivation.

- For some components it only makes sense to have one in the system (database repository, object factory).
- E.g., the construction call is expensive. We only do it once. We give everyone the same instance.
- Want to prevent anyone creating additional copies.
- Need to take care of lazy instantiation.

Summary.

- Lazy one-time initialization using *sync.Once*
- Adhere to Dependency Inversion Principle: depend on interfaces, not concrete types
- Singleton is not scary :-)

#### Singleton

```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sync"
)

type singletonDatabase struct {
  capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
  return db.capitals[name]
}

// both init and sync.Once are thread-safe but only sync.Once is lazy

var once sync.Once
var instance *singletonDatabase

func readData(path string) (map[string]int, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  result := map[string]int{}

  for scanner.Scan() {
    k := scanner.Text()
    scanner.Scan()
    v, _ := strconv.Atoi(scanner.Text())
    result[k] = v
  }

  return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
  once.Do(func() {
    db := singletonDatabase{}
    caps, err := readData("capitals.txt")
    if err == nil {
      db.capitals = caps
    }
    instance = &db
  })
  return instance
}

func main() {
  db := GetSingletonDatabase()
  pop := db.GetPopulation("Seoul")
  fmt.Println("Pop of Seoul =", pop)
}
```

#### Problems with Singleton

In our implementation we depend on real data, so unit testing becomes problematic.

This also breaks the Dependency Inversion Principle. We should depend on interface instead.

```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sync"
)

type singletonDatabase struct {
  capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
  return db.capitals[name]
}

var once sync.Once
var instance *singletonDatabase

func readData(path string) (map[string]int, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  result := map[string]int{}

  for scanner.Scan() {
    k := scanner.Text()
    scanner.Scan()
    v, _ := strconv.Atoi(scanner.Text())
    result[k] = v
  }

  return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
  once.Do(func() {
    db := singletonDatabase{}
    caps, err := readData("capitals.txt")
    if err == nil {
      db.capitals = caps
    }
    instance = &db
  })
  return instance
}

func GetTotalPopulation(cities []string) int {
  result := 0
  for _, city := range cities {
    result += GetSingletonDatabase().GetPopulation(city)
  } //        ^^^^^^^^^^^^^^^^^^^^ breaks Dependency Inversion Principle
  return result
}

func main() {
  cities := []string{"Seoul", "Mexico City"}
  tp := GetTotalPopulation(cities)
  ok := tp == (17500000 + 17400000)
  fmt.Println(ok)
}
```

#### Singleton and Dependency Inversion

```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sync"
)

// think of a module as a singleton
type Database interface {
  GetPopulation(name string) int
}

type singletonDatabase struct {
  capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(
  name string) int {
  return db.capitals[name]
}

var once sync.Once
var instance Database

func readData(path string) (map[string]int, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  result := map[string]int{}

  for scanner.Scan() {
    k := scanner.Text()
    scanner.Scan()
    v, _ := strconv.Atoi(scanner.Text())
    result[k] = v
  }

  return result, nil
}

func GetSingletonDatabase() Database {
  once.Do(func() {
    db := singletonDatabase{}
    caps, err := readData("capitals.txt")
    if err == nil {
      db.capitals = caps
    }
    instance = &db
  })
  return instance
}

func GetTotalPopulationEx(db Database, cities []string) int {
  result := 0
  for _, city := range cities {
    result += db.GetPopulation(city)
  }
  return result
}

type DummyDatabase struct {
  dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
  if len(d.dummyData) == 0 {
    d.dummyData = map[string]int{
      "alpha": 1,
      "beta":  2,
      "gamma": 3}
  }
  return d.dummyData[name]
}

func main() {
  db := GetSingletonDatabase()
  pop := db.GetPopulation("Seoul")
  fmt.Println("Pop of Seoul = ", pop)

  cities := []string{"Seoul", "Mexico City"}
  tp := GetTotalPopulationEx(GetSingletonDatabase(), cities)
  ok := tp == (17500000 + 17400000) // testing on live data
  fmt.Println(ok)

  names := []string{"alpha", "gamma"} // expect 4
  tp = GetTotalPopulationEx(&DummyDatabase{}, names)
  ok = tp == 4
  fmt.Println(ok)
}
```

## Structural Design Patterns

### Adapter -- getting the interface you want from the interface you have

- Electrical devices have different power (interface) requirements (voltage: 5V or 2020V, socket type: Europe or UK or US).
- We cannot modify our gadgets to support every possible interface. Some support possible, e.g. 120/220V.
- Thus, we use a special device (an adapter) to give us the interface we require from the interface we have.

A construct which adapts an existing interface X to conform to the required interface Y.

- Implementing an Adapter is easy.
- Determine the API you have and the API you need.
- Create a component which aggregates (has a pointer to, ...) the adaptee.
- Intermediate representations can pile up: use caching and other optimizations.

#### Adapter

```go
package main

import (
  "fmt"
  "strings"
)

func minmax (a, b int) (int, int) {
  if a < b {
    return a, b
  } else {
    return b, a
  }
}

// ↑↑↑ utility functions

type Line struct {
  X1, Y1, X2, Y2 int
}

type VectorImage struct {
  Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
  width -= 1
  height -= 1
  return &VectorImage{[]Line {
    {0, 0, width, 0},
    {0, 0, 0, height},
    {width, 0, width, height},
    {0, height, width, height}}}
}

// ↑↑↑ the interface you're given
// ↓↓↓ the interface you have

type Point struct {
  X, Y int
}

type RasterImage interface {
  GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
  maxX, maxY := 0, 0
  points := owner.GetPoints()
  for _, pixel := range points {
    if pixel.X > maxX { maxX = pixel.X }
    if pixel.Y > maxY { maxY = pixel.Y }
  }
  maxX += 1
  maxY += 1

  // preallocate

  data := make([][]rune, maxY)
  for i := 0; i < maxY; i++ {
    data[i] = make([]rune, maxX)
    for j := range data[i] { data[i][j] = ' ' }
  }

  for _, point := range points {
    data[point.Y][point.X] = '*'
  }

  b := strings.Builder{}
  for _, line := range data {
    b.WriteString(string(line))
    b.WriteRune('\n')
  }

  return b.String()
}

// problem: I want to print a RasterImage
//          but I can only make a VectorImage

type vectorToRasterAdapter struct {
  points []Point
}

func (a *vectorToRasterAdapter) addLine(line Line) {
  left, right := minmax(line.X1, line.X2)
  top, bottom := minmax(line.Y1, line.Y2)
  dx := right - left
  dy := line.Y2 - line.Y1

  if dx == 0 {
    for y := top; y <= bottom; y++ {
      a.points = append(a.points, Point{left, y})
    }
  } else if dy == 0 {
    for x := left; x <= right; x++ {
      a.points = append(a.points, Point{x, top})
    }
  }

  fmt.Println("generated", len(a.points), "points")
}

func (a vectorToRasterAdapter) GetPoints() []Point {
  return a.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
  adapter := vectorToRasterAdapter{}
  for _, line := range vi.Lines {
    adapter.addLine(line)
  }

  return adapter // as RasterImage
}

func main() {
  rc := NewRectangle(6, 4)
  a := VectorToRaster(rc) // adapter!
  _ = VectorToRaster(rc)  // adapter!
  fmt.Print(DrawPoints(a))
}
```

#### Adapter Caching

```go
package main

import (
  "crypto/md5"
  "encoding/json"
  "fmt"
  "strings"
)

func minmax (a, b int) (int, int) {
  if a < b {
    return a, b
  } else {
    return b, a
  }
}

// ↑↑↑ utility functions

type Line struct {
  X1, Y1, X2, Y2 int
}

type VectorImage struct {
  Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
  width -= 1
  height -= 1
  return &VectorImage{[]Line {
    {0, 0, width, 0},
    {0, 0, 0, height},
    {width, 0, width, height},
    {0, height, width, height}}}
}

// ↑↑↑ the interface you're given
// ↓↓↓ the interface you have

type Point struct {
  X, Y int
}

type RasterImage interface {
  GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
  maxX, maxY := 0, 0
  points := owner.GetPoints()
  for _, pixel := range points {
    if pixel.X > maxX { maxX = pixel.X }
    if pixel.Y > maxY { maxY = pixel.Y }
  }
  maxX += 1
  maxY += 1

  // preallocate

  data := make([][]rune, maxY)
  for i := 0; i < maxY; i++ {
    data[i] = make([]rune, maxX)
    for j := range data[i] { data[i][j] = ' ' }
  }

  for _, point := range points {
    data[point.Y][point.X] = '*'
  }

  b := strings.Builder{}
  for _, line := range data {
    b.WriteString(string(line))
    b.WriteRune('\n')
  }

  return b.String()
}

// problem: I want to print a RasterImage
//          but I can only make a VectorImage

type vectorToRasterAdapter struct {
  points []Point
}

var pointCache = map[[16]byte] []Point{}

func (a *vectorToRasterAdapter) addLineCached(line Line) {
  hash := func (obj interface{}) [16]byte {
    bytes, _ := json.Marshal(obj)
    return md5.Sum(bytes)
  }
  h := hash(line)
  if pts, ok := pointCache[h]; ok {
    for _, pt := range pts {
      a.points = append(a.points, pt)
    }
    return
  }

  left, right := minmax(line.X1, line.X2)
  top, bottom := minmax(line.Y1, line.Y2)
  dx := right - left
  dy := line.Y2 - line.Y1

  if dx == 0 {
    for y := top; y <= bottom; y++ {
      a.points = append(a.points, Point{left, y})
    }
  } else if dy == 0 {
    for x := left; x <= right; x++ {
      a.points = append(a.points, Point{x, top})
    }
  }

  // be sure to add these to the cache
  pointCache[h] = a.points
  fmt.Println("generated", len(a.points), "points")
}

func (a vectorToRasterAdapter) GetPoints() []Point {
  return a.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
  adapter := vectorToRasterAdapter{}
  for _, line := range vi.Lines {
    adapter.addLineCached(line)
  }

  return adapter // as RasterImage
}

func main() {
  rc := NewRectangle(6, 4)
  a := VectorToRaster(rc) // adapter!
  _ = VectorToRaster(rc)  // adapter!
  fmt.Print(DrawPoints(a))
}
```

### Bridge -- connecting components together through abstractiona

- Bridge prevents a 'Cartesian product' complexity explosion.
- Example
  - Common type ThreadScheduler
  - Can be preemptive or cooperative
  - Can run on Windows or Unix
  - End up with 2x2 scenario: WindowsPts, UnixPts, WndowsCts, UnixCts
- Bridge pattern avoids the entity explosion.

Instead of having large single tree you have several smaller flat trees.

Bridge is a mechanism that decouples an interface (hierarchy) from an implementation (hierarchy).

- Decouple abstraction from implementation
- Both can exist as hierarchies
- A stronger form of encapsulation

#### Bridge

```go
package main

import "fmt"

type Renderer interface {
  RenderCircle(radius float32)
}

type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
  fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
  Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
  fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
  renderer Renderer
  radius float32
}

func (c *Circle) Draw() {
  c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float32) *Circle {
  return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Resize(factor float32) {
  c.radius *= factor
}

func main() {
  vector := VectorRenderer{}
  circle := NewCircle(&vector, 5)
  circle.Draw()

  raster := RasterRenderer{}
  circle = NewCircle(&raster, 5)
  circle.Draw()
}
```

### Composite -- treating individual and aggregate objects uniformly

- Objecs use other objects' fields / methods through embedding
- Composition lets us make compound objects
  - E.g., a mathematical expression composed of simple expressions; or
  - A shape group made of several different shapes
- Composite design pattern is used to treat both single (scalar) and composite objects uniformly
  - I.e., Foo and []Foo have common APIs

A mechanism for treating for treating individual (scalar) objects and compositions of objects in a uniform manner.

- Objects can use other objects via composition
- Some composed and singular objects need similar / identical behaviors
- Composite design patern lets us treat both types of objects uniformly
- Iteration supported with the Iterator design pattern

#### Geometric Shapes

```go
package main

import (
  "fmt"
  "strings"
)

type GraphicObject struct {
  Name, Color string
  Children []GraphicObject
}

func (g *GraphicObject) String() string {
  sb := strings.Builder{}
  g.print(&sb, 0)
  return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
  sb.WriteString(strings.Repeat("*", depth))
  if len(g.Color) > 0 {
    sb.WriteString(g.Color)
    sb.WriteRune(' ')
  }
  sb.WriteString(g.Name)
  sb.WriteRune('\n')

  for _, child := range g.Children {
    child.print(sb, depth+1)
  }
}

func NewCircle(color string) *GraphicObject {
  return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
  return &GraphicObject{"Square", color, nil}
}

func main() {
  drawing := GraphicObject{"My Drawing", "", nil}
  drawing.Children = append(drawing.Children, *NewSquare("Red"))
  drawing.Children = append(drawing.Children, *NewCircle("Yellow"))

  group := GraphicObject{"Group 1", "", nil}
  group.Children = append(group.Children, *NewCircle("Blue"))
  group.Children = append(group.Children, *NewSquare("Blue"))
  drawing.Children = append(drawing.Children, group)

  fmt.Println(drawing.String())
}
```

#### Neural Networks

```go
package main

type NeuronInterface interface {
  Iter() []*Neuron
}

type Neuron struct {
  In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
  return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
  n.Out = append(n.Out, other)
  other.In = append(other.In, n)
}

type NeuronLayer struct {
  Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
  result := make([]*Neuron, 0)
  for i := range n.Neurons {
    result = append(result, &n.Neurons[i])
  }
  return result
}

func NewNeuronLayer(count int) *NeuronLayer {
  return &NeuronLayer{make([]Neuron, count)}
}

func Connect(left, right NeuronInterface) {
  for _, l := range left.Iter() {
    for _, r := range right.Iter() {
      l.ConnectTo(r)
    }
  }
}

func main() {
  neuron1, neuron2 := &Neuron{}, &Neuron{}
  layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

  //neuron1.ConnectTo(&neuron2)

  Connect(neuron1, neuron2)
  Connect(neuron1, layer1)
  Connect(layer2, neuron1)
  Connect(layer1, layer2)
}
```

### Decorator -- Adding behaviour without altering the type itself

- Want to augment an object with additional functionality
- Do not want to rewrite or alter existing code (follow Open Close Principle, open for extension closed for modification)
- Want to keep new functionality separate (follow Separation of Responsibilities Principle)
- Need to be able to interact with existing structures

Solution: embed the decorated object and provide additional functionality.

Decorator facilitates the addition of behaviors to individual objects through embedding.

- A decorator embeds the decorated objects
- Adds utility fields and methods to augment the object's features
- Often used to emulate multiple inheritance (may require extra work)

#### Multiple Inheritance

Let's construct a Dragon which is a Bird and a Lizard at the same time.

```go
package main

import "fmt"

/*
type Bird struct {
  Age int
}

func (b *Bird) Fly() {
  if b.Age >= 10 {
    fmt.Println("Flying!")
  }
}

type Lizard struct {
  Age int
}

func (l *Lizard) Crawl() {
  if l.Age < 10 {
    fmt.Println("Crawling!")
  }
}

type Dragon struct {
  Bird
  Lizard
}
*/

type Aged interface {
  Age() int
  SetAge(age int)
}

type Bird struct {
  age int
}

func (b *Bird) Age() int { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
  if b.age >= 10 {
    fmt.Println("Flying!")
  }
}

type Lizard struct {
  age int
}

func (l *Lizard) Age() int { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
  if l.age < 10 {
    fmt.Println("Crawling!")
  }
}

type Dragon struct {
  bird Bird
  lizard Lizard
}

func (d *Dragon) Age() int {
  return d.bird.age
}

func (d *Dragon) SetAge(age int) {
  d.bird.SetAge(age)
  d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
  d.bird.Fly()
}

func (d *Dragon) Crawl() {
  d.lizard.Crawl()
}

func NewDragon() *Dragon {
  return &Dragon{Bird{}, Lizard{}}
}

func main() {
  //d := Dragon{}
  //d.Bird.Age = 10
  //fmt.Println(d.Lizard.Age)
  //d.Fly()
  //d.Crawl()

  d := NewDragon()
  d.SetAge(10)
  d.Fly()
  d.Crawl()
}
```

#### Decorator

```go
package main

import "fmt"

type Shape interface {
  Render() string
}

type Circle struct {
  Radius float32
}

func (c *Circle) Render() string {
  return fmt.Sprintf("Circle of radius %f",
    c.Radius)
}

func (c *Circle) Resize(factor float32) {
  c.Radius *= factor
}

type Square struct {
  Side float32
}

func (s *Square) Render() string {
  return fmt.Sprintf("Square with side %f", s.Side)
}

// possible, but not generic enough
type ColoredSquare struct {
  Square
  Color string
}

type ColoredShape struct {
  Shape Shape
  Color string
}

func (c *ColoredShape) Render() string {
  return fmt.Sprintf("%s has the color %s",
    c.Shape.Render(), c.Color)
}

type TransparentShape struct {
  Shape Shape
  Transparency float32
}

func (t *TransparentShape) Render() string {
  return fmt.Sprintf("%s has %f%% transparency",
    t.Shape.Render(), t.Transparency * 100.0)
}

func main() {
  circle := Circle{2}
  fmt.Println(circle.Render())

  redCircle := ColoredShape{&circle, "Red"}
  fmt.Println(redCircle.Render())

  rhsCircle := TransparentShape{&redCircle, 0.5}
  fmt.Println(rhsCircle.Render())
}
```

### Façade -- Exposing several components through a single interface

- Balancing complexity and presentation / usability
- Typical home
  - Many subsystems (electrical, sanitation)
  - Complex internal structure (e.g., floor layers)
  - End user is not exposed to internals
- Same with software!
  - Many systems working to provide flexibility, but ...
  - API consumers want it to 'just work'

Façade provides a simple, easy to understand user interface over a large and sophisticated body of code.

- Build a Façade to provide a simplified API over a set of components
- May wish to (optionally) expose internals through the façade
- May allow users to 'escalate' to use more complex API if they need to

#### Façade

```go
package main

type Buffer struct {
  width, height int
  buffer []rune
}

func NewBuffer(width, height int) *Buffer {
  return &Buffer { width, height,
    make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
  return b.buffer[index]
}

type Viewport struct {
  buffer *Buffer
  offset int
}

func NewViewport(buffer *Buffer) *Viewport {
  return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
  return v.buffer.At(v.offset + index)
}

// a façade over buffers and viewports
type Console struct {
  buffers []*Buffer
  viewports []*Viewport
  offset int
}

func NewConsole() *Console {
  b := NewBuffer(10, 10)
  v := NewViewport(b)
  return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
  return c.viewports[0].GetCharacterAt(index)
}

func main() {
  c := NewConsole()
  u := c.GetCharacterAt(1)
}
```

### Flyweight -- Space optimization

- Avoid redundancy when storing data
- E.g., MMORPG
  - Plenty of users with identical first/last names
  - No sense in storing same first/last name over and over again
  - Store a list of names and references to them (indices, pointers, etc.)
- E.g., bold or italics text formatting
  - Don't want each character to have formatting character
  - Operate on *ranges* (e.g., line number, start/end positions)

Flyweight is a space optimization technique that lets us use less memory by storing externally the data associated with similar objects.

- Store common data externally
- Specify an index or a pointer into the external data store
- Define the idea of 'ranges' on homogenous collections and store data related to those ranges

#### Text Formatting

```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

type FormattedText struct {
  plainText  string
  capitalize []bool
}

func (f *FormattedText) String() string {
  sb := strings.Builder{}
  for i := 0; i < len(f.plainText); i++ {
    c := f.plainText[i]
    if f.capitalize[i] {
      sb.WriteRune(unicode.ToUpper(rune(c)))
    } else {
      sb.WriteRune(rune(c))
    }
  }
  return sb.String()
}

func NewFormattedText(plainText string) *FormattedText {
  return &FormattedText{plainText,
    make([]bool, len(plainText))}
}

func (f *FormattedText) Capitalize(start, end int) {
  for i := start; i <= end; i++ {
    f.capitalize[i] = true
  }
}

type TextRange struct {
  Start, End int
  Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
  return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
  plainText string
  formatting []*TextRange
}

func (b *BetterFormattedText) String() string {
  sb := strings.Builder{}

  for i := 0; i < len(b.plainText); i++ {
    c := b.plainText[i]
    for _, r := range b.formatting {
      if r.Covers(i) && r.Capitalize {
        c = uint8(unicode.ToUpper(rune(c)))
      }
    }
    sb.WriteRune(rune(c))
  }

  return sb.String()
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
  return &BetterFormattedText{plainText: plainText}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
  r := &TextRange{start, end, false, false, false}
  b.formatting = append(b.formatting, r)
  return r
}

func main() {
  text := "This is a brave new world"

  ft := NewFormattedText(text)
  ft.Capitalize(10, 15) // brave
  fmt.Println(ft.String())

  bft := NewBetterFormattedText(text)
  bft.Range(16, 19).Capitalize = true // new
  fmt.Println(bft.String())
}
```

#### User Names

```go
package main

import (
  "fmt"
  "strings"
)

type User struct {
  FullName string
}

func NewUser(fullName string) *User {
  return &User{FullName: fullName}
}

var allNames []string
type User2 struct {
  names []uint8
}

func NewUser2(fullName string) *User2 {
  getOrAdd := func(s string) uint8 {
    for i := range allNames {
      if allNames[i] == s {
        return uint8(i)
      }
    }
    allNames = append(allNames, s)
    return uint8(len(allNames) - 1)
  }

  result := User2{}
  parts := strings.Split(fullName, " ")
  for _, p := range parts {
    result.names = append(result.names, getOrAdd(p))
  }
  return &result
}

func (u *User2) FullName() string {
  var parts []string
  for _, id := range u.names {
    parts = append(parts, allNames[id])
  }
  return strings.Join(parts, " ")
}

func main() {
  john := NewUser("John Doe")
  jane := NewUser("Jane Doe")
  alsoJane := NewUser("Jane Smith")
  fmt.Println(john.FullName)
  fmt.Println(jane.FullName)
  fmt.Println("Memory taken by users:",
    len([]byte(john.FullName)) +
      len([]byte(alsoJane.FullName)) +
      len([]byte(jane.FullName)))

  john2 := NewUser2("John Doe")
  jane2 := NewUser2("Jane Doe")
  alsoJane2 := NewUser2("Jane Smith")
  fmt.Println(john2.FullName())
  fmt.Println(jane2.FullName())
  totalMem := 0
  for _, a := range allNames {
    totalMem += len([]byte(a))
  }
  totalMem += len(john2.names)
  totalMem += len(jane2.names)
  totalMem += len(alsoJane2.names)
  fmt.Println("Memory taken by users2:", totalMem)
}
```

### Proxy -- An interface for accessing a particular resource

- You are calling `foo.Bar()`
- This assumes that `foo` is in the same process as `Bar()`
- What if, later on, you want to put all `Foo`-related operations into a separate process?
  - Can you avoid changing your code?
- Proxy to the resque!
  - Same interface, entirely different behavior
- This is called a `communication proxy`
  - Other types: logging, virtual, guarding, ...

Proxy is a type that functions as an interface to a particular resource. That resource may be remote, expensive to construct, or may require logging or some other functionality.

- A proxy has the same interface as the underlying object
- To create a proxy, simply replicate the existing interface of an object
- Add relevant functionality to the redefined methods
- Different proxies (communication, logging, caching, etc.) have completely different behaviors

#### Proxy vs. Decorator

- Proxy tries to provide an identical interface; decorator provides an enhanced interface
- Decorator typically aggregates (or has pointer to) what it is decorating; proxy doesn't have to
- Proxy might not even be working with a materialized object

#### Protection Proxy

```go
package main

import "fmt"

type Driven interface {
  Drive()
}

type Car struct {}

func (c *Car) Drive() {
  fmt.Println("Car being driven")
}

type Driver struct {
  Age int
}

type CarProxy struct {
  car Car
  driver *Driver
}

func (c *CarProxy) Drive() {
  if c.driver.Age >= 16 {
    c.car.Drive()
  } else {
    fmt.Println("Driver too young")
  }
}

func NewCarProxy(driver *Driver) *CarProxy {
  return &CarProxy{Car{}, driver}
}

func main() {
  car := NewCarProxy(&Driver{12})
  car.Drive()
}
```

#### Virtual Proxy

```go
package main

import "fmt"

type Image interface {
  Draw()
}

type Bitmap struct {
  filename string
}

func (b *Bitmap) Draw() {
  fmt.Println("Drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
  fmt.Println("Loading image from", filename)
  return &Bitmap{filename: filename}
}

func DrawImage(image Image) {
  fmt.Println("About to draw the image")
  image.Draw()
  fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
  filename string
  bitmap *Bitmap
}

func (l *LazyBitmap) Draw() {
  if l.bitmap == nil {
    l.bitmap = NewBitmap(l.filename)
  }
  l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
  return &LazyBitmap{filename: filename}
}

func main() {
  //bmp := NewBitmap("demo.png")
  bmp := NewLazyBitmap("demo.png")
  DrawImage(bmp)
}
```

## Behavioral Design Patterns

### Chain of Responsibility -- Sequence of handlers processing an event one after another

- Unethical behavior by an employee; who takes the blame?
  - Employee
  - Manager
  - CEO
- You click a graphical element on a form
  - Button handles it, stops further processing
  - Underlying group box
  - Underlying window
- CCG computer game
  - Creature has attack and defense values
  - Those can be boosted by other cards

A chain of components who all get a chance to process a command or a query, optionally having default processing implementation and an ability to terminate the processing chain.

- Chain of Responsibility can be implemented as a linked list of pointers or a centralized construct
- Enlist objects in the chain, possibky controlling their order
- Control object removal from chain

#### Method Chain

```go
package chainofresponsibility

import "fmt"

type Creature struct {
  Name string
  Attack, Defense int
}

func (c *Creature) String() string {
  return fmt.Sprintf("%s (%d/%d)",
    c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
  return &Creature{Name: name, Attack: attack, Defense: defense}
}

type Modifier interface {
  Add(m Modifier)
  Handle()
}

type CreatureModifier struct {
  creature *Creature
  next Modifier // singly linked list
}

func (c *CreatureModifier) Add(m Modifier) {
  if c.next != nil {
    c.next.Add(m)
  } else { c.next = m }
}

func (c *CreatureModifier) Handle() {
  if c.next != nil {
    c.next.Handle()
  }
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
  return &CreatureModifier{creature: creature}
}

type DoubleAttackModifier struct {
  CreatureModifier
}

func NewDoubleAttackModifier(
  c *Creature) *DoubleAttackModifier {
  return &DoubleAttackModifier{CreatureModifier{
    creature: c }}
}

type IncreasedDefenseModifier struct {
  CreatureModifier
}

func NewIncreasedDefenseModifier(
  c *Creature) *IncreasedDefenseModifier {
  return &IncreasedDefenseModifier{CreatureModifier{
    creature: c }}
}

func (i *IncreasedDefenseModifier) Handle() {
  if i.creature.Attack <= 2 {
    fmt.Println("Increasing",
      i.creature.Name, "\b's defense")
    i.creature.Defense++
  }
  i.CreatureModifier.Handle()
}

func (d *DoubleAttackModifier) Handle() {
  fmt.Println("Doubling", d.creature.Name,
    "attack...")
  d.creature.Attack *= 2
  d.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
  CreatureModifier
}

func NewNoBonusesModifier(
  c *Creature) *NoBonusesModifier {
  return &NoBonusesModifier{CreatureModifier{
    creature: c }}
}

func (n *NoBonusesModifier) Handle() {
  // nothing here!
}

func main() {
  goblin := NewCreature("Goblin", 1, 1)
  fmt.Println(goblin.String())

  root := NewCreatureModifier(goblin)

  //root.Add(NewNoBonusesModifier(goblin))
  root.Add(NewDoubleAttackModifier(goblin))
  root.Add(NewIncreasedDefenseModifier(goblin))
  root.Add(NewDoubleAttackModifier(goblin))

  // eventually process the entire chain
  root.Handle()
  fmt.Println(goblin.String())
}
```

#### Command Query Separation

- Command = asking for an action or change (e.g., please set your attack value to 2).
- Query = asking for information (e.g., please give me your attack value).
- CQS = having separate means of sending commands and queries to e.g., direct field access.

#### Broker Chain

```go
package main

import (
  "fmt"
  "sync"
)

// cqs, mediator, cor

type Argument int

const (
  Attack Argument = iota
  Defense
)

type Query struct {
  CreatureName string
  WhatToQuery Argument
  Value int
}

type Observer interface {
  Handle(*Query)
}

type Observable interface {
  Subscribe(o Observer)
  Unsubscribe(o Observer)
  Fire(q *Query)
}

type Game struct {
  observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
  g.observers.Store(o, struct{}{})
  //                   ↑↑↑ empty anon struct
}

func (g *Game) Unsubscribe(o Observer) {
  g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
  g.observers.Range(func(key, value interface{}) bool {
    if key == nil {
      return false
    }
    key.(Observer).Handle(q)
    return true
  })
}

type Creature struct {
  game *Game
  Name string
  attack, defense int // ← private!
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
  return &Creature{game: game, Name: name, attack: attack, defense: defense}
}

func (c *Creature) Attack() int {
  q := Query{c.Name, Attack, c.attack}
  c.game.Fire(&q)
  return q.Value
}

func (c *Creature) Defense() int {
  q := Query{c.Name, Defense, c.defense}
  c.game.Fire(&q)
  return q.Value
}

func (c *Creature) String() string {
  return fmt.Sprintf("%s (%d/%d)",
    c.Name, c.Attack(), c.Defense())
}

// data common to all modifiers
type CreatureModifier struct {
  game *Game
  creature *Creature
}

func (c *CreatureModifier) Handle(*Query) {
  // nothing here!
}

type DoubleAttackModifier struct {
  CreatureModifier
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
  d := &DoubleAttackModifier{CreatureModifier{g, c}}
  g.Subscribe(d)
  return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
  if q.CreatureName == d.creature.Name &&
    q.WhatToQuery == Attack {
    q.Value *= 2
  }
}

func (d *DoubleAttackModifier) Close() error {
  d.game.Unsubscribe(d)
  return nil
}

func main() {
  game := &Game{sync.Map{}}
  goblin := NewCreature(game, "Strong Goblin", 2, 2)
  fmt.Println(goblin.String())

  {
    m := NewDoubleAttackModifier(game, goblin)
    fmt.Println(goblin.String())
    m.Close()
  }

  fmt.Println(goblin.String())
}
```

### Command -- You shall not pass

- Ordinary statements are perishable
  - Cannot undo field assignment
  - Cannot directly serialize a sequence of actions (calls)
- Want an object that represents an operation
  - `person` should change its `age` to 22
  - `car` should `explode()`
- Uses: GUI commands, multi-level undo/redo, macro recording and more

Command is an object which represents an instruction to perform a particular action. Contains all the information necessary for the action to be taken.

- Encapsulate all details of an operation in a separate object
- Define functions for applying the command (either in the command itself or elsewhere)
- Optionally define instructions for undoing the command
- Can create composite commands (a.k.a. macros)

#### Command

```go
package command

import "fmt"

var overdraftLimit = -500
type BankAccount struct {
  balance int
}

func (b *BankAccount) Deposit(amount int) {
  b.balance += amount
  fmt.Println("Deposited", amount,
    "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) {
  if b.balance - amount >= overdraftLimit {
    b.balance -= amount
    fmt.Println("Withdrew", amount,
      "\b, balance is now", b.balance)
  }
}

type Command interface {
  Call()
  Undo()
}

type Action int
const (
  Deposit Action = iota
  Withdraw
)

type BankAccountCommand struct {
  account *BankAccount
  action Action
  amount int
}

func (b *BankAccountCommand) Call() {
  switch b.action {
  case Deposit:
    b.account.Deposit(b.amount)
  case Withdraw:
    b.account.Withdraw(b.amount)
  }
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
  return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
  ba := BankAccount{}
  cmd := NewBankAccountCommand(&ba, Deposit, 100)
  cmd.Call()
  cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
  cmd2.Call()
  fmt.Println(ba)
}
```

#### Undo Operations

```go
package command

import "fmt"

var overdraftLimit = -500
type BankAccount struct {
  balance int
}

func (b *BankAccount) Deposit(amount int) {
  b.balance += amount
  fmt.Println("Deposited", amount,
    "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
  if b.balance - amount >= overdraftLimit {
    b.balance -= amount
    fmt.Println("Withdrew", amount,
      "\b, balance is now", b.balance)
    return true
  }
  return false
}

type Command interface {
  Call()
  Undo()
}

type Action int
const (
  Deposit Action = iota
  Withdraw
)

type BankAccountCommand struct {
  account *BankAccount
  action Action
  amount int
  succeeded bool
}

func (b *BankAccountCommand) Call() {
  switch b.action {
  case Deposit:
    b.account.Deposit(b.amount)
    b.succeeded = true
  case Withdraw:
    b.succeeded = b.account.Withdraw(b.amount)
  }
}

func (b *BankAccountCommand) Undo() {
  if !b.succeeded { return }
  switch b.action {
  case Deposit:
    b.account.Withdraw(b.amount)
  case Withdraw:
    b.account.Deposit(b.amount)
  }
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
  return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
  ba := BankAccount{}
  cmd := NewBankAccountCommand(&ba, Deposit, 100)
  cmd.Call()
  cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
  cmd2.Call()
  fmt.Println(ba)
}
```

#### Composite Command

```go
package main

import "fmt"

var overdraftLimit = -500
type BankAccount struct {
  balance int
}

func (b *BankAccount) Deposit(amount int) {
  b.balance += amount
  fmt.Println("Deposited", amount,
    "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
  if b.balance - amount >= overdraftLimit {
    b.balance -= amount
    fmt.Println("Withdrew", amount,
      "\b, balance is now", b.balance)
    return true
  }
  return false
}

type Command interface {
  Call()
  Undo()
  Succeeded() bool
  SetSucceeded(value bool)
}

type Action int
const (
  Deposit Action = iota
  Withdraw
)

type BankAccountCommand struct {
  account *BankAccount
  action Action
  amount int
  succeeded bool
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
  b.succeeded = value
}

// additional member
func (b *BankAccountCommand) Succeeded() bool {
  return b.succeeded
}

func (b *BankAccountCommand) Call() {
  switch b.action {
  case Deposit:
    b.account.Deposit(b.amount)
    b.succeeded = true
  case Withdraw:
    b.succeeded = b.account.Withdraw(b.amount)
  }
}

func (b *BankAccountCommand) Undo() {
  if !b.succeeded { return }
  switch b.action {
  case Deposit:
    b.account.Withdraw(b.amount)
  case Withdraw:
    b.account.Deposit(b.amount)
  }
}

type CompositeBankAccountCommand struct {
  commands []Command
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
  for _, cmd := range c.commands {
    if !cmd.Succeeded() {
      return false
    }
  }
  return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(value bool) {
  for _, cmd := range c.commands {
    cmd.SetSucceeded(value)
  }
}

func (c *CompositeBankAccountCommand) Call() {
  for _, cmd := range c.commands {
    cmd.Call()
  }
}

func (c *CompositeBankAccountCommand) Undo() {
  // undo in reverse order
  for idx := range c.commands {
    c.commands[len(c.commands)-idx-1].Undo()
  }
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
  return &BankAccountCommand{account: account, action: action, amount: amount}
}

type MoneyTransferCommand struct {
  CompositeBankAccountCommand
  from, to *BankAccount
  amount int
}

func NewMoneyTransferCommand(from *BankAccount, to *BankAccount, amount int) *MoneyTransferCommand {
  c := &MoneyTransferCommand{from: from, to: to, amount: amount}
  c.commands = append(c.commands,
    NewBankAccountCommand(from, Withdraw, amount))
  c.commands = append(c.commands,
    NewBankAccountCommand(to, Deposit, amount))
  return c
}

func (m *MoneyTransferCommand) Call() {
  ok := true
  for _, cmd := range m.commands {
    if ok {
      cmd.Call()
      ok = cmd.Succeeded()
    } else {
      cmd.SetSucceeded(false)
    }
  }
}

func main() {
  ba := &BankAccount{}
  cmdDeposit := NewBankAccountCommand(ba, Deposit, 100)
  cmdWithdraw := NewBankAccountCommand(ba, Withdraw, 1000)
  cmdDeposit.Call()
  cmdWithdraw.Call()
  fmt.Println(ba)
  cmdWithdraw.Undo()
  cmdDeposit.Undo()
  fmt.Println(ba)

  from := BankAccount{100}
  to := BankAccount{0}
  mtc := NewMoneyTransferCommand(&from, &to, 100) // try 1000
  mtc.Call()

  fmt.Println("from=", from, "to=", to)

  fmt.Println("Undoing...")
  mtc.Undo()
  fmt.Println("from=", from, "to=", to)
}
```

#### Functional Command

```go
package main

import "fmt"

type BankAccount struct {
  Balance int
}

func Deposit(ba *BankAccount, amount int) {
  fmt.Println("Depositing", amount)
  ba.Balance += amount
}

func Withdraw(ba *BankAccount, amount int) {
  if ba.Balance >= amount {
    fmt.Println("Withdrawing", amount)
    ba.Balance -= amount
  }
}

func main() {
  ba := &BankAccount{0}
  var commands []func()
  commands = append(commands, func() {
    Deposit(ba, 100)
  })
  commands = append(commands, func() {
    Withdraw(ba, 100)
  })

  for _, cmd := range commands {
    cmd()
  }
}
```

### Interpreter -- Interpreters are all arount us. Even now, in this very room

- Textual input needs to be processed
  - E.g., turned into linked structures
  - AST - Abstract Syntax Tree
- Some examples
  - Programming language compilers, interpreters and IDEs
  - HTML, XML and similar
  - Numeric expressions (3+4/5)
  - Regular expressions
- Turning strings into linked structure is a complicated process

An interpreter is a component that processes structured text data. Does so by turning it into separate lexical tokens (lexing) and then interpreting sequences of said tokens (parsing).

- Barring simple cases, an interpreter acts in two stages
- Lexing turns text into a set of tokens, e.g.
  3*(4+5) -> Lit[3] Star Lparen Lit[4] Plus Lit[5] Rparen
- Parsing tokens into meaningful constructs
  (AST = Abstract Syntax Tree)
  -> MultiplicationExpression[Integer[3], AdditionExpression[Integer[4], Integer[5]]]
- Parsed data can be then traversed using the Visitor pattern

#### Lexing

```go
package main

import (
  "fmt"
  "strconv"
  "strings"
  "unicode"
)

type TokenType int

const (
  Int TokenType = iota
  Plus
  Minus
  Lparen
  Rparen
)

type Token struct {
  Type TokenType
  Text string
}

func (t *Token) String() string {
  return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
  var result []Token

  // not using range here
  for i := 0; i < len(input); i++ {
    switch input[i] {
    case '+':
      result = append(result, Token{Plus, "+"})
    case '-':
      result = append(result, Token{Minus, "-"})
    case '(':
      result = append(result, Token{Lparen, "("})
    case ')':
      result = append(result, Token{Rparen, ")"})
    default:
      sb := strings.Builder{}
      for j := i; j < len(input); j++ {
        if unicode.IsDigit(rune(input[j])) {
          sb.WriteRune(rune(input[j]))
          i++
        } else {
          result = append(result, Token{
            Int, sb.String() })
          i--
          break
        }
      }
    }
  }
  return result
}

func main() {
  input := "(13+4)-(12+1)"
  tokens := Lex(input)
  fmt.Println(tokens)
}
```

#### Parsing

```go
package main

import (
  "fmt"
  "strconv"
  "strings"
  "unicode"
)

type Element interface {
  Value() int
}

type Integer struct {
  value int
}

func NewInteger(value int) *Integer {
  return &Integer{value: value}
}

func (i *Integer) Value() int {
  return i.value
}

type Operation int

const (
  Addition Operation = iota
  Subtraction
)

type BinaryOperation struct {
  Type Operation
  Left, Right Element
}

func (b *BinaryOperation) Value() int {
  switch b.Type {
  case Addition:
    return b.Left.Value() + b.Right.Value()
  case Subtraction:
    return b.Left.Value() + b.Right.Value()
  default:
    panic("Unsupported operation")
  }
}

type TokenType int

const (
  Int TokenType = iota
  Plus
  Minus
  Lparen
  Rparen
)

type Token struct {
  Type TokenType
  Text string
}

func (t *Token) String() string {
  return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
  var result []Token

  // not using range here
  for i := 0; i < len(input); i++ {
    switch input[i] {
    case '+':
      result = append(result, Token{Plus, "+"})
    case '-':
      result = append(result, Token{Minus, "-"})
    case '(':
      result = append(result, Token{Lparen, "("})
    case ')':
      result = append(result, Token{Rparen, ")"})
    default:
      sb := strings.Builder{}
      for j := i; j < len(input); j++ {
        if unicode.IsDigit(rune(input[j])) {
          sb.WriteRune(rune(input[j]))
          i++
        } else {
          result = append(result, Token{
            Int, sb.String() })
          i--
          break
        }
      }
    }
  }
  return result
}

func Parse(tokens []Token) Element {
  result := BinaryOperation{}
  haveLhs := false
  for i := 0; i < len(tokens); i++ {
    token := &tokens[i]
    switch token.Type {
    case Int:
      n, _ := strconv.Atoi(token.Text)
      integer := Integer{n}
      if !haveLhs {
        result.Left = &integer
        haveLhs = true
      } else {
        result.Right = &integer
      }
    case Plus:
      result.Type = Addition
    case Minus:
      result.Type = Subtraction
    case Lparen:
      j := i
      for ; j < len(tokens); j++ {
        if tokens[j].Type == Rparen {
          break
        }
      }
      // now j points to closing bracket, so
      // process subexpression without opening
      var subexp []Token
      for k := i+1; k < j; k++ {
        subexp = append(subexp, tokens[k])
      }
      element := Parse(subexp)
      if !haveLhs {
        result.Left = element
        haveLhs = true
      } else {
        result.Right = element
      }
      i = j
    }
  }
  return &result
}

func main() {
  input := "(13+4)-(12+1)"
  tokens := Lex(input)
  fmt.Println(tokens)

  parsed := Parse(tokens)
  fmt.Printf("%s = %d\n",
    input, parsed.Value())
}
```

### Iterator -- How traversal of data structures happens and who makes it happen

- Iteration (traversal) is a core functionality of various data structures
- An `iterator` is a type that facilitates the traversal
  - Keeps a pointer to the current element
  - Knows how to move to a different element
- Go allows iteration with `range`
  - Built-in support in many objects (arrays, slices, etc.)
  - Can be supported in our own structs

An object that facilitates the traversal of a data structure.

- An iterator specifies how you can traverse an object
- Moves along the iterated collection, indicating when last element has been reached
- Not idiomatic in Go (no standard Iterable interface)

#### Iteration

There are three ways of iteration in Go: range (keyword), generator (channels), explicit iterator (not idiomatic, almost the same as in C++).

```go
package main

import "fmt"

type Person struct {
  FirstName, MiddleName, LastName string
}

func (p *Person) Names() []string {
  return []string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
  out := make(chan string)
  go func() {
    defer close(out)
    out <- p.FirstName
    if len(p.MiddleName) > 0 {
      out <- p.MiddleName
    }
    out <- p.LastName
  }()
  return out
}

type PersonNameIterator struct {
  person *Person
  current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
  return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
  p.current++
  return p.current < 3
}

func (p *PersonNameIterator) Value() string {
  switch p.current {
    case 0: return p.person.FirstName
    case 1: return p.person.MiddleName
    case 2: return p.person.LastName
  }
  panic("We should not be here!")
}

func main() {
  p := Person{"Alexander", "Graham", "Bell"}

  // range
  for _, name := range p.Names() {
    fmt.Println(name)
  }

  // generator
  for name := range p.NamesGenerator() {
    fmt.Println(name)
  }

  // iterator
  for it := NewPersonNameIterator(&p); it.MoveNext(); {
    fmt.Println(it.Value())
  }
}
```

#### Tree Traversal

```go
package iterator

import "fmt"

type Node struct {
  Value int
  left, right, parent *Node
}

func NewNode(value int, left *Node, right *Node) *Node {
  n := &Node{Value: value, left: left, right: right}
  left.parent = n
  right.parent = n
  return n
}

func NewTerminalNode(value int) *Node {
  return &Node{Value:value}
}

type InOrderIterator struct {
  Current *Node
  root *Node
  returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
  i := &InOrderIterator{
    Current:       root,
    root:          root,
    returnedStart: false,
  }
  // move to the leftmost element
  for ;i.Current.left != nil; {
    i.Current = i.Current.left
  }
  return i
}

func (i *InOrderIterator) Reset() {
  i.Current = i.root
  i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
  if i.Current == nil { return false }
  if !i.returnedStart {
    i.returnedStart = true
    return true // can use first element
  }

  if i.Current.right != nil {
    i.Current = i.Current.right
    for ;i.Current.left != nil; {
      i.Current = i.Current.left
    }
    return true
  } else {
    p := i.Current.parent
    for ;p != nil && i.Current == p.right; {
      i.Current = p
      p = p.parent
    }
    i.Current = p
    return i.Current != nil
  }
}

type BinaryTree struct {
  root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
  return &BinaryTree{root: root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
  return NewInOrderIterator(b.root)
}

func main() {
  //   1
  //  / \
  // 2   3

  // in-order:  213
  // preorder:  123
  // postorder: 231

  root := NewNode(1,
    NewTerminalNode(2),
    NewTerminalNode(3))
  it := NewInOrderIterator(root)

  for ;it.MoveNext(); {
    fmt.Printf("%d,", it.Current.Value)
  }
  fmt.Println("\b")

  t := NewBinaryTree(root)
  for i := t.InOrder(); i.MoveNext(); {
    fmt.Printf("%d,", i.Current.Value)
  }
  fmt.Println("\b")
}
```

### Mediator -- Facilitates communication between components

- Components may go in and out of a system at any time
  - Chat room participants
  - Players in an MMORPG
- It makes no sense for them to have direct references (pointers) to one another
  - Those references may go dead
- Solution: have them all refer to some central component that facilitates communication

A component that facilitates communication between other components without them necessarily being aware of each other or having direct (reference) access to each other.

- Create the mediator and have each object in the system point to it
  - E.g., assign a field in factory function
- Mediator engages in bidirectional communication with its connected components
- Mediator has methods the components can call
- Components have methods the mediator can call
- Event processing (e.g., Rx) libraries make communication easier to implement

#### Chat Room

```go
package mediator

import "fmt"

type Person struct {
  Name string
  Room *ChatRoom
  chatLog []string
}

func NewPerson(name string) *Person {
  return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
  s := fmt.Sprintf("%s: '%s'", sender, message)
  fmt.Printf("[%s's chat session] %s\n", p.Name, s)
  p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(message string) {
  p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
  p.Room.Message(p.Name, who, message)
}

type ChatRoom struct {
  people []*Person
}

func (c *ChatRoom) Broadcast(source, message string) {
  for _, p := range c.people {
    if p.Name != source {
      p.Receive(source, message)
    }
  }
}

func (c *ChatRoom) Join(p *Person) {
  joinMsg := p.Name + " joins the chat"
  c.Broadcast("Room", joinMsg)

  p.Room = c
  c.people = append(c.people, p)
}

func (c *ChatRoom) Message(src, dst, msg string) {
  for _, p := range c.people {
    if p.Name == dst {
      p.Receive(src, msg)
    }
  }
}

func main() {
  room := ChatRoom{}

  john := NewPerson("John")
  jane := NewPerson("Jane")

  room.Join(john)
  room.Join(jane)

  john.Say("hi room")
  jane.Say("oh, hey john")

  simon := NewPerson("Simon")
  room.Join(simon)
  simon.Say("hi everyone!")

  jane.PrivateMessage("Simon", "glad you could join us!")
}
```

### Memento -- Keep a memento of an object's state to return to that state

- An object or system goes through changes
  - E.g., bank account gets deposits and withdrawals
- There are different ways of navigating those changes
- One way is to record every change (Command) and teach a command to 'undo' itself
  - Also part of CQRS = Command Query Responsibility Segregation
- Another is to simply save snapshots of the system

Memento is a token representing the system state. Let us roll back to the state when the token was generated. May or may not directly expose state information.

- Mementos are used to roll back states arbitrarily
- A memento is simply a token/handle with (typically) no methods of its own
- A memento is not required to expose directly the state(s) to which it reverts the system
- Can be used to implement undo/redo

#### Memento vs Flyweight

- Both patterns provide a 'token' clients can hold on to
- Memento is used *only* to be fed back into system
  - No public/mutable state
  - No methods
- A flyweight is similar to an ordinary reference to object
  - Can mutate state
  - Can provide additional functionality (fields/methods)

#### Memento

```go
package memento

import (
  "fmt"
)

type Memento struct {
  Balance int
}

type BankAccount struct {
  balance int
}

func (b *BankAccount) Deposit(amount int) *Memento {
  b.balance += amount
  return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
  b.balance = m.Balance
}

func main() {
  ba := BankAccount{100}
  m1 := ba.Deposit(50)
  m2 := ba.Deposit(25)
  fmt.Println(ba)

  ba.Restore(m1)
  fmt.Println(ba) // 150

  ba.Restore(m2)
  fmt.Println(ba)
}
```

#### Undo and Redo

```go
package memento

import "fmt"

type Memento struct {
  Balance int
}

type BankAccount struct {
  balance int
  changes []*Memento
  current int
}

func (b *BankAccount) String() string {
  return fmt.Sprint("Balance = $", b.balance,
    ", current = ", b.current)
}

func NewBankAccount(balance int) *BankAccount {
  b := &BankAccount{balance: balance}
  b.changes = append(b.changes, &Memento{balance})
  return b
}

func (b *BankAccount) Deposit(amount int) *Memento {
  b.balance += amount
  m := Memento{b.balance}
  b.changes = append(b.changes, &m)
  b.current++
  fmt.Println("Deposited", amount,
    ", balance is now", b.balance)
  return &m
}

func (b *BankAccount) Restore(m *Memento) {
  if m != nil {
    b.balance -= m.Balance
    b.changes = append(b.changes, m)
    b.current = len(b.changes) - 1
  }
}

func (b *BankAccount) Undo() *Memento {
  if b.current > 0 {
    b.current--
    m := b.changes[b.current]
    b.balance = m.Balance
    return m
  }
  return nil // nothing to undo
}

func (b *BankAccount) Redo() *Memento {
  if b.current + 1 < len(b.changes) {
    b.current++
    m := b.changes[b.current]
    b.balance = m.Balance
    return m
  }
  return nil
}

func main() {
  ba := NewBankAccount(100)
  ba.Deposit(50)
  ba.Deposit(25)
  fmt.Println(ba)

  ba.Undo()
  fmt.Println("Undo 1:", ba)
  ba.Undo()
  fmt.Println("Undo 2:", ba)
  ba.Redo()
  fmt.Println("Redo:", ba)
}
```

### Observer -- I am watching you

- We need to be informed when certain things happen
  - Object's field changes
  - Object does something
  - Some external event occurs
- We want to listen to events and be notified when they occur
- Two participants: *observable* and *observer*

An *observer* is an object that wishes to be informed about events happening in the system. The entity generating the events is an *observable*.

No standard implementation in Go.

- Observer is an intrusive approach
- Must provide a way of clients to subscribe
- Event data sent from observable to all subscribers
- Data represented as *interface{}*
- Unsubscription is possible

#### Observer and Observable

```go
package observer

import (
  "container/list"
  "fmt"
)

type Observable struct {
  subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
  o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    if z.Value.(Observer) == x {
      o.subs.Remove(z)
    }
  }
}

func (o *Observable) Fire(data interface{}) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    z.Value.(Observer).Notify(data)
  }
}

type Observer interface {
  Notify(data interface{})
}

// whenever a person catches a cold,
// a doctor must be called
type Person struct {
  Observable
  Name string
}

func NewPerson(name string) *Person {
  return &Person {
    Observable: Observable{new(list.List)},
    Name: name,
  }
}

func (p *Person) CatchACold() {
  p.Fire(p.Name)
}

type DoctorService struct {}

func (d *DoctorService) Notify(data interface{}) {
  fmt.Printf("A doctor has been called for %s",
    data.(string))
}

func main() {
  p := NewPerson("Boris")
  ds := &DoctorService{}
  p.Subscribe(ds)

  // let's test it!
  p.CatchACold()
}
```

#### Property Observers

```go
package observer

import (
  "container/list"
  "fmt"
)

type Observable struct {
  subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
  o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    if z.Value.(Observer) == x {
      o.subs.Remove(z)
    }
  }
}

func (o *Observable) Fire(data interface{}) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    z.Value.(Observer).Notify(data)
  }
}

type Observer interface {
  Notify(data interface{})
}

type Person struct {
  Observable
  age int
}

func NewPerson(age int) *Person {
  return &Person{Observable{new(list.List)}, age}
}

type PropertyChanged struct {
  Name string
  Value interface{}
}

func (p *Person) Age() int { return p.age }
func (p *Person) SetAge(age int) {
  if age == p.age { return } // no change
  p.age = age
  p.Fire(PropertyChanged{"Age", p.age})
}

type TrafficManagement struct {
  o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
  if pc, ok := data.(PropertyChanged); ok {
    if pc.Value.(int) >= 16 {
      fmt.Println("Congrats, you can drive now!")
      // we no longer care
      t.o.Unsubscribe(t)
    }
  }
}

func main() {
  p := NewPerson(15)
  t := &TrafficManagement{p.Observable}
  p.Subscribe(t)

  for i := 16; i <= 20; i++ {
    fmt.Println("Setting age to", i)
    p.SetAge(i)
  }
}
```

#### Property Dependencies

```go
package observer

import (
  "container/list"
  "fmt"
)

type Observable struct {
  subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
  o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    if z.Value.(Observer) == x {
      o.subs.Remove(z)
    }
  }
}

func (o *Observable) Fire(data interface{}) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    z.Value.(Observer).Notify(data)
  }
}

type Observer interface {
  Notify(data interface{})
}

type Person struct {
  Observable
  age int
}

func NewPerson(age int) *Person {
  return &Person{Observable{new(list.List)}, age}
}

type PropertyChanged struct {
  Name string
  Value interface{}
}

func (p *Person) Age() int { return p.age }
func (p *Person) SetAge(age int) {
  if age == p.age { return } // no change

  oldCanVote := p.CanVote()

  p.age = age
  p.Fire(PropertyChanged{"Age", p.age})

  if oldCanVote != p.CanVote() {
    p.Fire(PropertyChanged{"CanVote", p.CanVote()})
  }
}

func (p *Person) CanVote() bool {
  return p.age >= 18
}

type ElectrocalRoll struct {
}

func (e *ElectrocalRoll) Notify(data interface{}) {
  if pc, ok := data.(PropertyChanged); ok {
    if pc.Name == "CanVote" && pc.Value.(bool) {
      fmt.Println("Congratulations, you can vote!")
    }
  }
}

func main() {
  p := NewPerson(0)
  er := &ElectrocalRoll{}
  p.Subscribe(er)

  for i := 10; i < 20; i++ {
    fmt.Println("Setting age to", i)
    p.SetAge(i)
  }
}
```

### State -- Fun with Finite State Machines

- Consider an ordinary telephone
- What do you do with it depends on the state of the phone/line
  - If it's ringing or you want to make a call, you can pick it up
  - Phone must be off the hook to talk/make a call
  - If you try calling someone, and it's busy, you put the handset down
- Changes in state can be explicit or in response to event (Observer pattern)

A pattern in which the object's behavior is determined by its stare. An object transitions from one state to another (something needs to *trigger* a transition).

A formalized construct which manages state and transitions is called a *state machine*.

- Given sufficient complexity, it pays to formally define possible states and events/triggers
- Can define
  - State entry/exit behaviors
  - Action when a particular event causes a transition
  - Guard conditions enabling/disabling transition
  - Default action when no transitions are found for an event

#### Classic Implementation

```go
package state

import "fmt"

type Switch struct {
  State State
}

func NewSwitch() *Switch {
  return &Switch{NewOffState()}
}

func (s *Switch) On() {
  s.State.On(s)
}

func (s *Switch) Off() {
  s.State.Off(s)
}

type State interface {
  On(sw *Switch)
  Off(sw *Switch)
}

type BaseState struct {}

func (s *BaseState) On(sw *Switch) {
  fmt.Println("Light is already on")
}

func (s *BaseState) Off(sw *Switch) {
  fmt.Println("Light is already off")
}

type OnState struct {
  BaseState
}

func NewOnState() *OnState {
  fmt.Println("Light turned on")
  return &OnState{BaseState{}}
}

func (o *OnState) Off(sw *Switch) {
  fmt.Println("Turning light off...")
  sw.State = NewOffState()
}

type OffState struct {
  BaseState
}

func NewOffState() *OffState {
  fmt.Println("Light turned off")
  return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
  fmt.Println("Turning light on...")
  sw.State = NewOnState()
}

func main() {
  sw := NewSwitch()
  sw.On()
  sw.Off()
  sw.Off()
}
```

#### Handmade State Machine

```go
package state

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

type State int

const (
  OffHook State = iota
  Connecting
  Connected
  OnHold
  OnHook
)

func (s State) String() string {
  switch s {
  case OffHook: return "OffHook"
  case Connecting: return "Connecting"
  case Connected: return "Connected"
  case OnHold: return "OnHold"
  case OnHook: return "OnHook"
  }
  return "Unknown"
}

type Trigger int

const (
  CallDialed Trigger = iota
  HungUp
  CallConnected
  PlacedOnHold
  TakenOffHold
  LeftMessage
)

func (t Trigger) String() string {
  switch t {
  case CallDialed: return "CallDialed"
  case HungUp: return "HungUp"
  case CallConnected: return "CallConnected"
  case PlacedOnHold: return "PlacedOnHold"
  case TakenOffHold: return "TakenOffHold"
  case LeftMessage: return "LeftMessage"
  }
  return "Unknown"
}

type TriggerResult struct {
  Trigger Trigger
  State State
}

var rules = map[State][]TriggerResult {
  OffHook: {
    {CallDialed, Connecting},
  },
  Connecting: {
    {HungUp, OffHook},
    {CallConnected, Connected},
  },
  Connected: {
    {LeftMessage, OnHook},
    {HungUp, OnHook},
    {PlacedOnHold, OnHold},
  },
  OnHold: {
    {TakenOffHold, Connected},
    {HungUp, OnHook},
  },
}

func main() {
  state, exitState := OffHook, OnHook
  for ok := true; ok; ok = state != exitState {
    fmt.Println("The phone is currently", state)
    fmt.Println("Select a trigger:")

    for i := 0; i < len(rules[state]); i++ {
      tr := rules[state][i]
      fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
    }

    input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
    i, _ := strconv.Atoi(string(input))

    tr := rules[state][i]
    state = tr.State
  }
  fmt.Println("We are done using the phone")
}
```

#### Switch-Based State Machine

```go
package state

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type State int

const (
  Locked State = iota
  Failed
  Unlocked
)

func main() {
  code := "1234"
  state := Locked
  entry := strings.Builder{}

  for {
    switch state {
    case Locked:
      // only reads input when you press Return
      r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
      entry.WriteRune(r)

      if entry.String() == code {
        state = Unlocked
        break
      }

      if strings.Index(code, entry.String()) != 0 {
        // code is wrong
        state = Failed
      }
    case Failed:
      fmt.Println("FAILED")
      entry.Reset()
      state = Locked
    case Unlocked:
      fmt.Println("UNLOCKED")
      return
    }
  }
}
```

### Strategy -- System behavior partially specified at runtime

- Many algorithms can be decomposed into higher- and lower-level parts
- Making tea can be decomposed into
  - The process of making a hot beverage (boil water, pour into cup); and
  - Tea-specific things (put teabag into water)
- The high-level algorithm can then be reused for making coffee or hot chocolate
  - Supported by beverage-specific strategies

 Separates an algorithm into its 'skeleton' and concrete implementation steps, which can be varied at run-time.

- Define an algorithm at a high level
- Define the interface you expect each strategy to follow
- Support the injection of the strategy into the high-level algorithm

#### Strategy

```go
package strategy

import (
  "fmt"
  "strings"
)

type OutputFormat int

const (
  Markdown OutputFormat = iota
  Html
)

type ListStrategy interface {
  Start(builder *strings.Builder)
  End(builder *strings.Builder)
  AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct {}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) AddListItem(
  builder *strings.Builder, item string) {
  builder.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct {}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
  builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
  builder.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
  builder.WriteString("  <li>" + item + "</li>\n")
}

type TextProcessor struct {
  builder strings.Builder
  listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
  return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
  switch fmt {
  case Markdown:
    t.listStrategy = &MarkdownListStrategy{}
  case Html:
    t.listStrategy = &HtmlListStrategy{}
  }
}

func (t *TextProcessor) AppendList(items []string) {
  t.listStrategy.Start(&t.builder)
  for _, item := range items {
    t.listStrategy.AddListItem(&t.builder, item)
  }
  t.listStrategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
  t.builder.Reset()
}

func (t *TextProcessor) String() string {
  return t.builder.String()
}

func main() {
  tp := NewTextProcessor(&MarkdownListStrategy{})
  tp.AppendList([]string{ "foo", "bar", "baz" })
  fmt.Println(tp)

  tp.Reset()
  tp.SetOutputFormat(Html)
  tp.AppendList([]string{ "foo", "bar", "baz" })
  fmt.Println(tp)
}
```

### Template Method -- A high-level blueprint for an algorithm to be completed by inheritors

- Algorithms can be decomposed into common parts + specifics
- Strategy pattern does this through composition
  - High-level algorithm uses an interface
  - Concrete implementations implement the interface
  - We keep a pointer to the interface; provide concrete implementations
- Template Method peforms a similar operation, but
  - It's typically just a function, not a struct with a reference to the implementation
  - Can still use interfaces (just like Strategy); or
  - Can be functional (take several functions as parameters)

A skeleton algorithm defined in a function. Function can either use an interface (like Strategy) or can take several functions as arguments.

- Very similar to Strategy
- Typical implementation:
  - Define an interface with common operations
  - Make use of those operations inside a function
- Alternative function approach:
  - Make a function that takes several functions
  - Can pass in functions that capture local state
  - No need for either structs or interfaces

#### Template Method

```go
package templatemethod

import "fmt"

type Game interface {
  Start()
  HaveWinner() bool
  TakeTurn()
  WinningPlayer() int
}

func PlayGame(g Game) {
  g.Start()
  for ;!g.HaveWinner(); {
    g.TakeTurn()
  }
  fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

type chess struct {
  turn, maxTurns, currentPlayer int
}

func NewGameOfChess() Game {
  return &chess{ 1, 10, 0 }
}

func (c *chess) Start() {
  fmt.Println("Starting a game of chess.")
}

func (c *chess) HaveWinner() bool {
  return c.turn == c.maxTurns
}

func (c *chess) TakeTurn() {
  c.turn++
  fmt.Printf("Turn %d taken by player %d\n",
    c.turn, c.currentPlayer)
  c.currentPlayer = (c.currentPlayer + 1) % 2
}

func (c *chess) WinningPlayer() int {
  return c.currentPlayer
}

func main() {
  chess := NewGameOfChess()
  PlayGame(chess)
}
```

#### Functional Template Method

```go
package templatemethod

import "fmt"

func PlayGame(start, takeTurn func(),
  haveWinner func()bool,
  winningPlayer func()int) {
  start()
  for ;!haveWinner(); {
    takeTurn()
  }
  fmt.Printf("Player %d wins.\n", winningPlayer())
}

func main() {
  turn, maxTurns, currentPlayer := 1, 10, 0

  start := func() {
    fmt.Println("Starting a game of chess.")
  }

  takeTurn := func() {
    turn++
    fmt.Printf("Turn %d taken by player %d\n",
      turn, currentPlayer)
    currentPlayer = (currentPlayer + 1) % 2
  }

  haveWinner := func()bool {
    return turn == maxTurns
  }

  winningPlayer := func()int {
    return currentPlayer
  }

  PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
```

### Visitor -- Allows adding extra behaviors to entire hierarchies of types

- Need to define a new operation on an entire type hierarchy
  - E.g., given a document model (lists, paragraphs, etc.), we want to add printing functionality
- Do not want to keep modifying every type in the hierarchy
- Want to have the new functionality separate (SRP = Separation of Responcibilities Principle)
- This approach is often used for traversal
  - Alternative to Iterator
  - Hierarchy members help you traverse themselves

A pattern where a component (visitor) is allowed to traverse the entire hierarchy of types. Implemented by propagating a single *Accept()* method throughout the entire hierarchy.

- Which function to call?
- Single dispatch: depends on name of request and type of receiver
- Double dispatch: depends on name of request and type of *two* receivers (type of visitor, type of element being visited)

- Propagate an `Accept(v *Visitor)` method throughout the entire hierarchy
- Create a visitor with `VisitFoo(f Foo)`, `VisitBar(b Bar)`, ... for each element in the hierarchy
- Each `Accept()` simply calls `Visitor.VisitXxx(self)`

#### Intrusive Visitor

Violates Open Closed Principle (Open for extension, Closed for modification).

```go
package visitor

import (
  "fmt"
  "strings"
)

type Expression interface {
  Print(sb *strings.Builder)
}

type DoubleExpression struct {
  value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
  sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
  left, right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
  sb.WriteString("(")
  a.left.Print(sb)
  sb.WriteString("+")
  a.right.Print(sb)
  sb.WriteString(")")
}

func main() {
  // 1+(2+3)
  e := AdditionExpression{
    &DoubleExpression{1},
    &AdditionExpression{
      left:  &DoubleExpression{2},
      right: &DoubleExpression{3},
    },
  }
  sb := strings.Builder{}
  e.Print(&sb)
  fmt.Println(sb.String())
}
```

#### Reflective Visitor

```go
package visitor

import (
  "fmt"
  "strings"
)

type Expression interface {
  // nothing here!
}

type DoubleExpression struct {
  value float64
}

type AdditionExpression struct {
  left, right Expression
}

func Print(e Expression, sb *strings.Builder) {
  if de, ok := e.(*DoubleExpression); ok {
    sb.WriteString(fmt.Sprintf("%g", de.value))
  } else if ae, ok := e.(*AdditionExpression); ok {
    sb.WriteString("(")
    Print(ae.left, sb)
    sb.WriteString("+")
    Print(ae.right, sb)
    sb.WriteString(")")
  }

  // breaks OCP
  // will work incorrectly on missing case
}

func main() {
  // 1+(2+3)
  e := &AdditionExpression{
    &DoubleExpression{1},
    &AdditionExpression{
      left:  &DoubleExpression{2},
      right: &DoubleExpression{3},
    },
  }
  sb := strings.Builder{}
  Print(e, &sb)
  fmt.Println(sb.String())
}
```

#### Classic Visitor

Classic double dispatch visitor.

```go
package visitor

import (
  "fmt"
  "strings"
)

type ExpressionVisitor interface {
  VisitDoubleExpression(de *DoubleExpression)
  VisitAdditionExpression(ae *AdditionExpression)
}

type Expression interface {
  Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
  value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
  ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
  left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
  ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
  sb strings.Builder
}

func (e *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
  e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
  e.sb.WriteString("(")
  ae.left.Accept(e)
  e.sb.WriteString("+")
  ae.right.Accept(e)
  e.sb.WriteString(")")
}

func NewExpressionPrinter() *ExpressionPrinter {
  return &ExpressionPrinter{strings.Builder{}}
}

func (e *ExpressionPrinter) String() string {
  return e.sb.String()
}

func main() {
  // 1+(2+3)
  e := &AdditionExpression{
    &DoubleExpression{1},
    &AdditionExpression{
      left:  &DoubleExpression{2},
      right: &DoubleExpression{3},
    },
  }
  ep := NewExpressionPrinter()
  ep.VisitAdditionExpression(e)
  fmt.Println(ep.String())
}
```

## Course Summary

Whew, that was a lot of patterns! What did we learn?

### Creational

- Builder
  - Separate component for when object construction gets too complicated
  - Can create mutually cooperative sub-builders
  - Often has a fluent interface
- Factories
  - Factory functions (constructors) are common
  - Factory can be a simple function or a dedicated struct
- Prototype
  - Creation of an object from an existing object
  - Requires either explicit deep copy or copy through serialization
- Singleton
  - When you need to ensure just a single instance exists
  - Can be made thread-safe and lazy
  - Consider extracting interface or using dependency injection

### Structural

- Adapter
  - Convert the interface you get to the interface you need
- Bridge
  - Decouple abstraction from implementation
- Composite
  - Allows clients to treat individual objects and compositions of objects uniformly
- Decorator
  - Attach additional responsibilities to objects
  - Can be done through embedding or pointers
- Façade
  - Provide a single unified interface over a set of interfaces
- Flyweight
  - Efficiently support very large number of similar objects
- Proxy
  - Provide a surrogate object that forwards calls to the real object while performing additional functions
  - E.g., access control, communication, logging etc.

### Behavioral

- Chain of Responsibility
  - Allow components to process information/events in a chain
  - Each element in the chain refers to next element; or
  - Make a list and go through it
- Command
  - Encapsulate a request into a separate object
  - Good for audit, replay, undo/redo
  - Part of CQS/CQRS (Command Query Separation / Command Query Responsibility Segregation)
- Interpreter
  - Transform textual input into structures (e.g. ASTs)
  - Used by interpreters, compilers, static analysis tools, etc.
  - *Compiler Theory* is a separate branch of Computer Science
_ Iterator
  - Provides an interface for accessing elements of an aggregate object
- Mediator
  - Provides mediation services between several objects
  - E.g., message passing, chat room
- Memento
  - Yields token representing system states
  - Tokens do not allow direct manipulation, but can be used in appropriate APIs
- Observer
  - Allows notifications of changes/happenings in a component
- State
  - We model systems by having one of a possible states and transitions between these states
  - Such a system is called a *state machine*
  - Special frameworks exist to orchestrate state machines
- Strategy & template Method
  - Both define a skeleton algorithm with details filled in by implementer
  - Strategy uses composition; Template Method doesn't
- Visitor
  - Allows non-intrusive addition of functionality to hierarchies
