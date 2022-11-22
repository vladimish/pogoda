package open_meteo

type Forecast struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	HourlyUnits          struct {
		Time          string `json:"time"`
		Temperature2M string `json:"temperature_2m"`
		Precipitation string `json:"precipitation"`
		Cloudcover    string `json:"cloudcover"`
	} `json:"hourly_units"`
	Hourly struct {
		Time          []int64   `json:"time"`
		Temperature2M []float64 `json:"temperature_2m"`
		Precipitation []float64 `json:"precipitation"`
		Cloudcover    []int     `json:"cloudcover"`
	} `json:"hourly"`
}
