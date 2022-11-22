package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/vladimish/pogoda/internal/models"
	"github.com/vladimish/pogoda/internal/presets"
)

func (c *Controller) Start(userId int64, firstName string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(userId, "Hello, I'm a bot!")

	err := c.User.Add(&models.User{
		TelegramId: userId,
		FirstName:  firstName,
	})
	if err != nil {
		logrus.Error(err)
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	keyboard := presets.GenerateStartKeyboard()
	msg.ReplyMarkup = keyboard

	return msg
}
