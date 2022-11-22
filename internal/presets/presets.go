package presets

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	InternalErr = "Ошибка, попробуйте позже"
	BackBtn     = "◀"
)

func GenerateStartKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	markup := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(TodayBtn),
			tgbotapi.NewKeyboardButton(TomorrowBtn),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(FavouritesBtn),
		),
	)

	return &markup
}

func GenerateCancelKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	markup := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(BackBtn),
		),
	)

	return &markup
}
