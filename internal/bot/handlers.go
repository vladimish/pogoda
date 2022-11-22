package bot

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/vladimish/pogoda/internal/presets"
)

func (b *Bot) handleUpdate(c api.Update) error {
	var err error
	switch {
	case c.Message.IsCommand():
		err = b.handleCommand(c.Message)

	case c.Message.Text != "" && !c.Message.IsCommand() && c.Message.Location == nil:
		err = b.handleMessage(c.Message)

	case c.Message.Location != nil:
		err = b.handleLocation(c.Message)

	default:
		err = b.handleUnknown(c.Message)
	}
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleCommand(c *api.Message) error {
	var msg api.MessageConfig

	switch c.Text {
	case "/start":
		msg = b.handlers.Start(c.Chat.ID, c.Chat.FirstName)
	}

	err := b.api.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleMessage(c *api.Message) error {
	var msg api.MessageConfig

	switch c.Text {
	// Menu
	case presets.FavouritesBtn:
		msg = b.handlers.OpenFavourites(c.Chat.ID)
	case presets.TodayBtn:
		msg = b.handlers.Today(c.Chat.ID)
	case presets.TomorrowBtn:
		msg = b.handlers.Tomorrow(c.Chat.ID)
	// ===

	// Favourites
	case presets.FavouritesAdd:
		msg = b.handlers.AddFavourites(c.Chat.ID)
	case presets.FavouritesDel:
		msg = b.handlers.DelFavourites(c.Chat.ID)
	// ===

	// Shared
	case presets.BackBtn:
		msg = b.handlers.Back(c.Chat.ID)

	default:
		msg = b.handlers.Unknown(c.Chat.ID, c.Text)
		// ===
	}

	err := b.api.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleLocation(c *api.Message) error {
	var msg api.MessageConfig

	msg = b.handlers.Location(c.Chat.ID, c.Location)

	err := b.api.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleUnknown(c *api.Message) error {
	logrus.Error("unexpected message type")
	err := b.api.Send(b.handlers.Unknown(c.Chat.ID, "Возможно, вы отправили стикер с котом."))
	if err != nil {
		return err
	}

	return nil
}
