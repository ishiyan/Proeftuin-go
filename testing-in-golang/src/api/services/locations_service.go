package services

import (
	"proeftuin/testing-in-golang/src/api/domain/locations"
	"proeftuin/testing-in-golang/src/api/utils/errors"
	"proeftuin/testing-in-golang/src/api/providers/locations_provider"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationsService locationsServiceInterface
)

func init() {
	LocationsService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
