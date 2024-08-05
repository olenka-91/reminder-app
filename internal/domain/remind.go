package domain

import (
	"errors"
	"time"
)

var (
	ErrRemindNotFound = errors.New("remind not found")
)

type Remind struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Msg        string    `json:"msg"`
	RemindDate time.Time `json:"remind_date"`
}

type UpdateRemindInput struct {
	Title      *string    `json:"title"`
	Msg        *string    `json:"msg"`
	RemindDate *time.Time `json:"remind_date"`
}
