package main

import "fmt"

type person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type personBuilder struct {
	person *person // needs to be inited
}

func newPersonBuilder() *personBuilder {
	return &personBuilder{&person{}}
}

func (pb *personBuilder) Build() *person {
	return pb.person
}

func (pb *personBuilder) Works() *personJobBuilder {
	return &personJobBuilder{*pb}
}

func (pb *personBuilder) Lives() *personAddressBuilder {
	return &personAddressBuilder{*pb}
}

type personJobBuilder struct {
	personBuilder
}

func (pjb *personJobBuilder) At(companyName string) *personJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *personJobBuilder) AsA(position string) *personJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *personJobBuilder) Earning(annualIncome int) *personJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

type personAddressBuilder struct {
	personBuilder
}

func (pab *personAddressBuilder) At(streetAddress string) *personAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

func (pab *personAddressBuilder) In(city string) *personAddressBuilder {
	pab.person.City = city
	return pab
}

func (pab *personAddressBuilder) WithPostcode(postcode string) *personAddressBuilder {
	pab.person.Postcode = postcode
	return pab
}

func main() {
	pb := newPersonBuilder()
	pb.
		Lives().At("123 London Road").In("London").WithPostcode("SW12BC").
		Works().At("Fabrikam").AsA("Programmer").Earning(123000)
	person := pb.Build()
	fmt.Println(*person)
}
