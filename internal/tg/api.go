package tg

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type API struct {
	token string

	api *tg.BotAPI
}

func NewAPI(token string) *API {
	return &API{
		token: token,
	}
}

func (a *API) Start() tg.UpdatesChannel {
	bot, err := tg.NewBotAPI(a.token)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	a.api = bot

	return a.api.GetUpdatesChan(tg.NewUpdate(0))
}

func (a *API) Send(config tg.MessageConfig) error {
	_, err := a.api.Send(config)
	return err
}
