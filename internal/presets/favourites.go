package presets

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/vladimish/pogoda/internal/models"
)

const (
	TodayBtn      = "Сегодня"
	TomorrowBtn   = "Завтра"
	FavouritesBtn = "Избранное 🌟"
	FavouritesAdd = "Добавить населенный пункт"
	FavouritesDel = "Удалить населенный пункт"
)

func GenerateFavouritesKeyboard(favourites []models.Favourite) *tgbotapi.ReplyKeyboardMarkup {
	rows := make([][]tgbotapi.KeyboardButton, 0, len(favourites))
	if favourites != nil && len(favourites) > 0 {
		for i := range favourites {
			name := favourites[i].Name
			if favourites[i].Selected {
				name = "⭐ " + name
			}
			rows = append(rows, tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(name),
			))
		}
	}

	markup := tgbotapi.NewReplyKeyboard()

	markup.Keyboard = append(markup.Keyboard, rows...)
	markup.Keyboard = append(markup.Keyboard,
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(FavouritesAdd),
			tgbotapi.NewKeyboardButton(FavouritesDel),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(BackBtn),
		),
	)

	return &markup
}

func GenerateRemoveFavouriteKeyboard(favourites []models.Favourite) *tgbotapi.ReplyKeyboardMarkup {
	rows := make([][]tgbotapi.KeyboardButton, 0, len(favourites))
	if favourites != nil && len(favourites) > 0 {
		for i := range favourites {
			rows = append(rows, tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(favourites[i].Name),
			))
		}
	}

	markup := tgbotapi.NewReplyKeyboard()

	markup.Keyboard = append(markup.Keyboard, rows...)
	markup.Keyboard = append(markup.Keyboard,
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(BackBtn),
		),
	)

	return &markup
}
