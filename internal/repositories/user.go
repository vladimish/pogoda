package repositories

import (
	"github.com/vladimish/pogoda/internal/models"
)

type User interface {
	GetById(id int64) (*models.User, error)
	GetByTelegramId(chatId int64) (*models.User, error)
	Add(user *models.User) error
	UpdateState(id int64, state models.UserState) error
}
