package main

import (
	"context"

	"github.com/vladimish/pogoda/internal/bot"
	"github.com/vladimish/pogoda/internal/cfg"
	"github.com/vladimish/pogoda/internal/handlers"
	"github.com/vladimish/pogoda/internal/pg"
	"github.com/vladimish/pogoda/internal/tg"
)

func main() {
	c := cfg.Get()
	api := tg.NewAPI(c.TelegramToken)
	db, err := pg.NewDB(c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Database)
	if err != nil {
		panic(err)
	}

	usersRepo := pg.NewUser(db)
	favouritesRepo := pg.NewFavourite(db)
	ctrl := handlers.NewController(usersRepo, favouritesRepo)

	b := bot.NewBot(api, ctrl)
	err = b.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
