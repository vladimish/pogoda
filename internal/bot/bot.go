package bot

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/vladimish/pogoda/internal/handlers"
)

type Bot struct {
	token string

	handlers *handlers.Controller
	api      API
}

func NewBot(api API, c *handlers.Controller) *Bot {
	return &Bot{
		handlers: c,
		api:      api,
	}
}

func (b *Bot) Run(ctx context.Context) error {
	c := b.api.Start()

	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)
			err = b.Run(ctx)
			if err != nil {
				logrus.Panic(err)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done")
		case update := <-c:
			err := b.handleUpdate(update)
			if err != nil {
				logrus.Panic(err)
			}
		}
	}
}
