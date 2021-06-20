package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type peep struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	james := peep{
		Name:  "Bond, James Bond",
		Motto: "Shaken, not stirred",
	}

	miss := peep{
		Name:  "Miss Moneypenny",
		Motto: "If love could kill",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	peeps := []peep{james, miss}
	cars := []car{f, c}

	data := struct {
		Wisdom    []peep
		Transport []car
	}{
		peeps,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
