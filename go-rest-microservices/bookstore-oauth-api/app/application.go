package app

import (
	"proeftuin/go-rest-microservices/bookstore-oauth-api/services/access_token"
	"proeftuin/go-rest-microservices/bookstore-oauth-api/repository/db"
	"proeftuin/go-rest-microservices/bookstore-oauth-api/http"
	"github.com/gin-gonic/gin"
	"proeftuin/go-rest-microservices/bookstore-oauth-api/repository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
