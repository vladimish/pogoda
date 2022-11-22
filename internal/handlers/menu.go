package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vladimish/pogoda/internal/presets"
)

func (c *Controller) Menu(userId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(userId, "menu")
	msg.ReplyMarkup = presets.GenerateStartKeyboard()
	return msg
}
