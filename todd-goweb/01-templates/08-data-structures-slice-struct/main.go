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

	peeps := []peep{james, miss}

	err := tpl.Execute(os.Stdout, peeps)
	if err != nil {
		log.Fatalln(err)
	}
}
