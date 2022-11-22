package handlers

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/vladimish/pogoda/internal/presets"
	open_meteo "github.com/vladimish/pogoda/pkg/open-meteo"
)

func (c *Controller) Today(id int64) tgbotapi.MessageConfig {
	u, err := c.User.GetByTelegramId(id)
	if err != nil {
		logrus.Errorf("can't get user: %v", err)
		return tgbotapi.NewMessage(id, "Internal error")
	}

	favourites, err := c.Favourites.GetByUserId(u.Id)
	if err != nil {
		logrus.Errorf("can't get user: %v", err)
		return tgbotapi.NewMessage(id, presets.InternalErr)
	}

	if favourites == nil || len(favourites) == 0 {
		return tgbotapi.NewMessage(id, "Добавьте локацию в избранное")
	}
	fav := favourites[0]
	for i := range favourites {
		if favourites[i].Selected {
			fav = favourites[i]
			break
		}
	}

	forecast, err := open_meteo.GetForecast(fav.Lat, fav.Lon, time.Now(), time.Now().AddDate(0, 0, 1))
	if err != nil {
		logrus.Errorf("can't get forecast: %v", err)
		return tgbotapi.NewMessage(id, "Сервис погоды недоступен...")
	}

	msg := tgbotapi.NewMessage(id, fmt.Sprintf("Погода в %s на сегодня:\n%s", fav.Name, forecastMessage(forecast)))
	return msg
}

func (c *Controller) Tomorrow(id int64) tgbotapi.MessageConfig {
	u, err := c.User.GetByTelegramId(id)
	if err != nil {
		logrus.Errorf("can't get user: %v", err)
		return tgbotapi.NewMessage(id, "Internal error")
	}

	favourites, err := c.Favourites.GetByUserId(u.Id)
	if err != nil {
		logrus.Errorf("can't get user: %v", err)
		return tgbotapi.NewMessage(id, "Internal error")
	}

	if favourites == nil || len(favourites) == 0 {
		return tgbotapi.NewMessage(id, "Добавьте локацию в избранное")
	}
	fav := favourites[0]
	for i := range favourites {
		if favourites[i].Selected {
			fav = favourites[i]
			break
		}
	}

	forecast, err := open_meteo.GetForecast(fav.Lat, fav.Lon, time.Now().AddDate(0, 0, 1), time.Now().AddDate(0, 0, 2))
	if err != nil {
		logrus.Errorf("can't get forecast: %v", err)
		return tgbotapi.NewMessage(id, "Сервис погоды недоступен...")
	}

	msg := tgbotapi.NewMessage(id, fmt.Sprintf("Погода в %s на завтра:\n%s", fav.Name, forecastMessage(forecast)))
	return msg
}

func forecastMessage(forecast *open_meteo.Forecast) string {
	sb := strings.Builder{}
	for i := range forecast.Hourly.Time {
		t := time.Unix(forecast.Hourly.Time[i], 0)
		sb.WriteString(fmt.Sprintf("%s: %.1f °C; %s\n",
			t.Format("15:04"),
			forecast.Hourly.Temperature2M[i],
			getWeatherEmoji(forecast.Hourly.Cloudcover[i], forecast.Hourly.Precipitation[i])),
		)
	}

	return sb.String()
}

func getWeatherEmoji(cloudcover int, precipitation float64) string {
	switch {
	case cloudcover >= 50 && precipitation > 0 && precipitation < 3:
		return "🌧"
	case cloudcover >= 50 && precipitation >= 3:
		return "⛈"
	case cloudcover < 50 && precipitation > 0:
		return "🌦"
	case precipitation == 0 && cloudcover > 50:
		return "☁"
	case precipitation == 0 && cloudcover > 20 && cloudcover < 50:
		return "🌤"
	case precipitation == 0 && cloudcover <= 20:
		return "☀"
	}

	return "☀"
}
