package main

import (
	"net/http"
	"proeftuin/todd-goweb/10-mongo-db/06-hands-on-2/controllers"
	"proeftuin/todd-goweb/10-mongo-db/06-hands-on-2/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
