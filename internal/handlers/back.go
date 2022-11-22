package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/vladimish/pogoda/internal/models"
	"github.com/vladimish/pogoda/internal/presets"
)

func (c *Controller) Back(userId int64) tgbotapi.MessageConfig {
	u, err := c.User.GetByTelegramId(userId)
	if err != nil {
		logrus.Errorf("can't get user: %v", err)
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	msg := c.newBackMessage(userId, u.State)

	err = c.User.UpdateState(userId, statesDirections[u.State])
	if err != nil {
		logrus.Errorf("can't update user state: %v", err)
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	return msg
}

var statesDirections = map[models.UserState]models.UserState{
	models.FavouriteMenu:   models.Menu,
	models.FavouriteAdd:    models.FavouriteMenu,
	models.FavouriteDelete: models.FavouriteMenu,
}

func (c *Controller) newBackMessage(userId int64, state models.UserState) tgbotapi.MessageConfig {
	newState := statesDirections[state]
	msg := tgbotapi.NewMessage(userId, "back")
	switch newState {
	case models.Menu:
		msg.ReplyMarkup = presets.GenerateStartKeyboard()
		return msg
	case models.FavouriteMenu:
		f, err := c.Favourites.GetByUserId(userId)
		if err != nil {
			msg.ReplyMarkup = presets.GenerateStartKeyboard()
		} else {
			msg.ReplyMarkup = presets.GenerateFavouritesKeyboard(f)
		}

		return msg
	default:
		logrus.Error("unexpected state")
		msg := tgbotapi.NewMessage(userId, presets.InternalErr)
		return msg
	}
}
