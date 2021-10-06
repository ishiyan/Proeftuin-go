package main

import (
	"fitbit/fitbit"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	api := fitbit.New("23BKXX", "03e904359d3e6cb2c3cb0666074a453c", "http://localhost:80/auth")

	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query()["code"][0]
		api.LoadAccessToken(code)
		fmt.Println(api.GetProfile())
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, "Visit: <a href=%q>%q</a>", api.AuthorizeURI, api.AuthorizeURI)
	})

	fmt.Println("Visit: " + api.AuthorizeURI)
	log.Fatal(http.ListenAndServe("localhost:80", mux))
}
