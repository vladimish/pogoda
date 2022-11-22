package bigdatacloud

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Geocode struct {
	Latitude                  float64 `json:"latitude"`
	Longitude                 float64 `json:"longitude"`
	Continent                 string  `json:"continent"`
	LookupSource              string  `json:"lookupSource"`
	ContinentCode             string  `json:"continentCode"`
	LocalityLanguageRequested string  `json:"localityLanguageRequested"`
	City                      string  `json:"city"`
	CountryName               string  `json:"countryName"`
	CountryCode               string  `json:"countryCode"`
	Postcode                  string  `json:"postcode"`
	PrincipalSubdivision      string  `json:"principalSubdivision"`
	PrincipalSubdivisionCode  string  `json:"principalSubdivisionCode"`
	PlusCode                  string  `json:"plusCode"`
	Locality                  string  `json:"locality"`
	LocalityInfo              struct {
		Administrative []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			IsoName     string `json:"isoName,omitempty"`
			Order       int    `json:"order"`
			AdminLevel  int    `json:"adminLevel"`
			IsoCode     string `json:"isoCode,omitempty"`
			WikidataId  string `json:"wikidataId"`
			GeonameId   int    `json:"geonameId,omitempty"`
		} `json:"administrative"`
		Informative []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Order       int    `json:"order"`
			IsoCode     string `json:"isoCode,omitempty"`
			WikidataId  string `json:"wikidataId,omitempty"`
			GeonameId   int    `json:"geonameId,omitempty"`
		} `json:"informative"`
	} `json:"localityInfo"`
}

type Locality string

const (
	LocalityEn Locality = "en"
	LocalityRu Locality = "ru"
)

func GetGeocode(lat, lon float64, loc Locality) (*Geocode, error) {
	url := fmt.Sprintf("https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=%f&longitude=%f&localityLanguage=%s", lat, lon, loc)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var geocode Geocode
	err = json.NewDecoder(resp.Body).Decode(&geocode)
	if err != nil {
		return nil, err
	}

	return &geocode, nil
}
