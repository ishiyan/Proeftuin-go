package main

import (
	"proeftuin/todd-goweb/10-mongo-db/03-rest-controllers/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
