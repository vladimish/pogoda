package pg

import (
	"database/sql"
	"fmt"

	"github.com/vladimish/pogoda/internal/models"
)

type Favourite struct {
	db *sql.DB
}

func NewFavourite(db *sql.DB) *Favourite {
	return &Favourite{
		db: db,
	}
}

func (f *Favourite) GetById(id int64) (*models.Favourite, error) {
	res := f.db.QueryRow(`SELECT id, user_id, name, lon, lat, selected FROM favourites WHERE id = $1`, id)
	if res.Err() != nil {
		return nil, fmt.Errorf("cannot get favourite: %w", res.Err())
	}

	var favourite models.Favourite
	err := res.Scan(&favourite.Id, &favourite.UserId, &favourite.Name, &favourite.Lon, &favourite.Lat, &favourite.Selected)
	if err != nil {
		return nil, fmt.Errorf("cannot scan favourite: %w", err)
	}

	return &favourite, nil
}

func (f *Favourite) GetByUserId(userId int64) ([]models.Favourite, error) {
	rows, err := f.db.Query("SELECT id, user_id, name, lon, lat, selected FROM favourites WHERE user_id = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("cannot get favourites: %w", err)
	}

	var favourites []models.Favourite
	for rows.Next() {
		var favourite models.Favourite
		err := rows.Scan(&favourite.Id, &favourite.UserId, &favourite.Name, &favourite.Lon, &favourite.Lat, &favourite.Selected)
		if err != nil {
			return nil, fmt.Errorf("cannot scan favourite: %w", err)
		}
		favourites = append(favourites, favourite)
	}

	return favourites, nil
}

func (f *Favourite) Add(favourite *models.Favourite) error {
	// TODO: remove
	_, err := f.db.Exec("UPDATE favourites SET selected = false WHERE user_id = $1", favourite.UserId)
	if err != nil {
		return fmt.Errorf("cannot update favourites: %w", err)
	}

	row := f.db.QueryRow(
		"INSERT INTO favourites (user_id, name, lon, lat, selected) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		favourite.UserId, favourite.Name, favourite.Lon, favourite.Lat, favourite.Selected)
	if row.Err() != nil {
		return fmt.Errorf("cannot add favourite: %w", row.Err())
	}

	var id int64
	err = row.Scan(&id)
	if err != nil {
		return fmt.Errorf("cannot scan favourite id: %w", err)
	}

	favourite.Id = id
	return nil
}

func (f *Favourite) Delete(userId int64, text string) error {
	_, err := f.db.Exec("DELETE FROM favourites USING users WHERE users.telegram_id = $1 AND favourites.name = $2", userId, text)
	if err != nil {
		return fmt.Errorf("cannot delete favourite: %w", err)
	}

	return nil
}
