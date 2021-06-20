package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "DOGGY, Doggy, doggy")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "KITTY, Kitty, kitty")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "That's me!")
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
