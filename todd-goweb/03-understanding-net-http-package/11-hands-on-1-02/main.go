package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "DOGGY, Doggy, doggy")
}

func cat(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "KITTY, Kitty, kitty")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "That's me!")
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
