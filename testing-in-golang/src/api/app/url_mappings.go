package app

import (
	"proeftuin/testing-in-golang/src/api/controllers"
)

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
