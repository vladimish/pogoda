package geocode

import (
	"github.com/vladimish/pogoda/pkg/bigdatacloud"
)

func GetCityByLocation(lat, lon float64) (string, error) {
	geocode, err := bigdatacloud.GetGeocode(lat, lon, bigdatacloud.LocalityRu)
	if err != nil {
		return "", err
	}
	city := geocode.City
	if city == "" {
		city = geocode.Locality
	}
	return city, nil
}
