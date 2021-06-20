package app

import (
	"github.com/gin-gonic/gin"
	"proeftuin/go-rest-microservices/bookstore-utils-go/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8082")
}
