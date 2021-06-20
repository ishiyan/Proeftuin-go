package main

import (
	"encoding/json"
	"net/http"
	"log"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", handleEncode)
	http.ListenAndServe(":8080", nil)
}

func handleEncode(w http.ResponseWriter, r *http.Request) {
	ps := []person {
		person{ First: "James"},
		person{ First: "Jenny"},
	}

	err := json.NewEncoder(w).Encode(ps)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}
