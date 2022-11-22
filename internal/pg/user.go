package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/vladimish/pogoda/internal/models"
)

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

type User struct {
	db *sql.DB
}

func (u *User) GetById(id int64) (*models.User, error) {
	row := u.db.QueryRow(`
		SELECT id, telegram_id, first_name, state, register_date, last_message_date FROM users WHERE id = $1`, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.Id, &user.TelegramId, &user.FirstName, &user.State, &user.RegisterDate, &user.LastMessageDate)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetByTelegramId(chatId int64) (*models.User, error) {
	row := u.db.QueryRow(`
		SELECT id, telegram_id, first_name, state, register_date, last_message_date FROM users WHERE telegram_id = $1`, chatId)
	if row.Err() != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.Id, &user.TelegramId, &user.FirstName, &user.State, &user.RegisterDate, &user.LastMessageDate)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Add(user *models.User) error {
	res := u.db.QueryRow(`
		INSERT INTO users (telegram_id, first_name, register_date, last_message_date) VALUES ($1, $2, NOW(), NOW()) ON CONFLICT (telegram_id) DO UPDATE SET state=0 RETURNING id`,
		user.TelegramId, user.FirstName)
	if res.Err() != nil {
		return fmt.Errorf("can't insert user: %w", res.Err())
	}

	err := res.Scan(&user.Id)
	if err != nil {
		return fmt.Errorf("can't read inserted id: %w", err)
	}

	return nil
}

func (u *User) UpdateState(telegramId int64, state models.UserState) error {
	_, err := u.db.Exec(`
		UPDATE users SET state = $1 WHERE telegram_id = $2`, state, telegramId)
	if err != nil {
		return fmt.Errorf("can't update user state: %w", err)
	}

	return nil
}
