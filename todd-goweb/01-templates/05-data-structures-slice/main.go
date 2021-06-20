package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	peeps := []string{"Bond, James Bond", "Miss Moneypenny", "Double Zero"}

	err := tpl.Execute(os.Stdout, peeps)
	if err != nil {
		log.Fatalln(err)
	}
}
