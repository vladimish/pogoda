package models

import (
	"time"
)

type UserState byte

const (
	New             UserState = 0
	Menu            UserState = 1
	FavouriteMenu   UserState = 2
	FavouriteAdd    UserState = 3
	FavouriteDelete UserState = 4
)

type User struct {
	Id              int64
	TelegramId      int64
	FirstName       string
	State           UserState
	RegisterDate    time.Time
	LastMessageDate time.Time
}
