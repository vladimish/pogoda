package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/vladimish/pogoda/internal/geocode"
	"github.com/vladimish/pogoda/internal/models"
	"github.com/vladimish/pogoda/internal/presets"
)

func (c *Controller) Location(id int64, location *tgbotapi.Location) tgbotapi.MessageConfig {
	u, err := c.User.GetByTelegramId(id)
	if err != nil {
		logrus.Errorf("can't get user: %v", err)
		return tgbotapi.NewMessage(id, "Internal error")
	}

	if u.State != models.FavouriteAdd {
		logrus.Errorf("unexpected state: %v", u.State)
		return tgbotapi.NewMessage(id, "Internal error")
	}

	name, err := geocode.GetCityByLocation(location.Latitude, location.Longitude)
	if err != nil {
		logrus.Errorf("can't get city by location: %v", err)
		name = "Локация"
	}

	err = c.Favourites.Add(&models.Favourite{
		UserId:   u.Id,
		Name:     name,
		Lon:      location.Longitude,
		Lat:      location.Latitude,
		Selected: true,
	})
	if err != nil {
		logrus.Errorf("can't add favourite: %v", err)
		return tgbotapi.NewMessage(id, "Internal error")
	}

	err = c.User.UpdateState(id, models.Menu)
	if err != nil {
		logrus.Errorf("can't update user state: %v", err)
		return tgbotapi.NewMessage(id, "Internal error")
	}

	msg := tgbotapi.NewMessage(id, "Локация добавлена в избранное")
	keyboard := presets.GenerateStartKeyboard()
	msg.ReplyMarkup = keyboard

	return msg
}
