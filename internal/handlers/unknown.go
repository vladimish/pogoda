package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/vladimish/pogoda/internal/models"
)

const (
	UnknownMsg = "Неизвестная команда"
)

func (c *Controller) Unknown(userId int64, text string) tgbotapi.MessageConfig {
	u, err := c.User.GetByTelegramId(userId)
	if err != nil {
		logrus.Errorf("failed to get user by telegram id: %v", err)
		return tgbotapi.NewMessage(userId, UnknownMsg)
	}

	switch u.State {
	case models.FavouriteDelete:
		return c.FavouriteDelete(userId, text)
	default:
		return tgbotapi.NewMessage(userId, UnknownMsg)
	}
}
