package bot

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type API interface {
	Start() tg.UpdatesChannel
	Send(config tg.MessageConfig) error
}
