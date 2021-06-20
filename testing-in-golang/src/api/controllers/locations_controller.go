package controllers

import (
	"github.com/gin-gonic/gin"
	"proeftuin/testing-in-golang/src/api/services"
	"net/http"
)

func GetCountry(c *gin.Context) {
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
