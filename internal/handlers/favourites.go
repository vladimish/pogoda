package handlers

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/vladimish/pogoda/internal/models"
	"github.com/vladimish/pogoda/internal/presets"
)

func (c *Controller) OpenFavourites(userId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(userId, "weather")
	err := c.User.UpdateState(userId, models.FavouriteMenu)
	if err != nil {
		msg.Text = presets.InternalErr
		return msg
	}

	u, err := c.User.GetByTelegramId(userId)
	if err != nil {
		msg.Text = presets.InternalErr
		return msg
	}

	favourites, err := c.Favourites.GetByUserId(u.Id)
	if err != nil {
		msg.Text = presets.InternalErr
		return msg
	}

	keyboard := presets.GenerateFavouritesKeyboard(favourites)
	msg.ReplyMarkup = keyboard

	return msg
}

func (c *Controller) AddFavourites(userId int64) tgbotapi.MessageConfig {
	err := c.User.UpdateState(userId, models.FavouriteAdd)
	if err != nil {
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	msg := tgbotapi.NewMessage(userId, "Отправьте геометку населенного пункта")
	keyboard := presets.GenerateCancelKeyboard()
	msg.ReplyMarkup = keyboard

	return msg
}

func (c *Controller) DelFavourites(userId int64) tgbotapi.MessageConfig {
	err := c.User.UpdateState(userId, models.FavouriteDelete)
	if err != nil {
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	u, err := c.User.GetByTelegramId(userId)
	if err != nil {
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	favourites, err := c.Favourites.GetByUserId(u.Id)
	if err != nil {
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	msg := tgbotapi.NewMessage(userId, "Выберите город для удаления")
	keyboard := presets.GenerateRemoveFavouriteKeyboard(favourites)
	msg.ReplyMarkup = keyboard
	return msg
}

func (c *Controller) FavouriteDelete(userId int64, text string) tgbotapi.MessageConfig {
	text = strings.ReplaceAll(text, "⭐ ", "")

	err := c.Favourites.Delete(userId, text)
	if err != nil {
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	err = c.User.UpdateState(userId, models.Menu)
	if err != nil {
		return tgbotapi.NewMessage(userId, presets.InternalErr)
	}

	msg := tgbotapi.NewMessage(userId, "Удалено")
	keyboard := presets.GenerateStartKeyboard()
	msg.ReplyMarkup = keyboard
	return msg
}
