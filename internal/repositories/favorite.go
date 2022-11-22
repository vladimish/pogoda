package repositories

import (
	"github.com/vladimish/pogoda/internal/models"
)

type Favourite interface {
	GetById(id int64) (*models.Favourite, error)
	GetByUserId(userId int64) ([]models.Favourite, error)
	Add(favourite *models.Favourite) error
	Delete(userId int64, text string) error
}
