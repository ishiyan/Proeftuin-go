package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/decode", decode)
	http.ListenAndServe(":8080", nil)
}

func decode(w http.ResponseWriter, r *http.Request) {

	var ps []person

	err := json.NewDecoder(r.Body).Decode(&ps)
	if err != nil {
		log.Println("Decoded bad data", err)
	}

	fmt.Println("Decoded", ps)
}
