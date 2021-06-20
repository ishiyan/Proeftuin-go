package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p := person {
		First: "Jenny",
	}

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {

}
