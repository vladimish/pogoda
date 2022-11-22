package open_meteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const addr = "https://api.open-meteo.com/v1/forecast"

func GetForecast(lon, lat float64, from, to time.Time) (*Forecast, error) {
	req, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse url: %w", err)
	}

	q := req.Query()
	q.Set("latitude", fmt.Sprintf("%f", lat))
	q.Set("longitude", fmt.Sprintf("%f", lon))
	q.Set("start_date", from.Format("2006-01-02"))
	q.Set("end_date", to.Format("2006-01-02"))
	q.Set("timeformat", "unixtime")
	q.Set("temperature_unit", "celsius")
	req.RawQuery = q.Encode()
	// Using strings concatenation instead of q.Set() because
	// api violates RFC 3986 and returns 400 Bad Request if
	// we are encoding reserved symbol "comma" as %2C
	req.RawQuery += "&hourly=temperature_2m,precipitation,cloudcover"

	resp, err := http.Get(req.String())
	if err != nil {
		return nil, fmt.Errorf("cannot get forecast: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot get forecast: %s", resp.Status)
	}

	var forecast Forecast
	err = json.NewDecoder(resp.Body).Decode(&forecast)
	if err != nil {
		return nil, fmt.Errorf("cannot decode forecast: %w", err)
	}

	for i := range forecast.Hourly.Temperature2M {
		forecast.Hourly.Temperature2M[i] = float64(5) / 9 * (forecast.Hourly.Temperature2M[i] - 32)
	}

	return &forecast, nil
}
