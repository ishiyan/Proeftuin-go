package app

import (
	"proeftuin/go-different-microservice-design-patterns/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:id", controllers.UsersController.Get)
	router.POST("/users", controllers.UsersController.Save)
}
