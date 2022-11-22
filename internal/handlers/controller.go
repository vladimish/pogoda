package handlers

import (
	"github.com/vladimish/pogoda/internal/repositories"
)

type Controller struct {
	User       repositories.User
	Favourites repositories.Favourite
}

func NewController(user repositories.User, favourites repositories.Favourite) *Controller {
	return &Controller{
		User:       user,
		Favourites: favourites,
	}
}
