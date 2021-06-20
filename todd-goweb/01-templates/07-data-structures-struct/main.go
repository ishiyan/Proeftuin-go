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
		Name:  "James Bond",
		Motto: "Shaken, not stirred",
	}

	err := tpl.Execute(os.Stdout, james)
	if err != nil {
		log.Fatalln(err)
	}
}
