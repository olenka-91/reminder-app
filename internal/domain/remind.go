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
type RemindUpdateInput struct {
	Title      *string    `json:"title"`
	Msg        *string    `json:"msg"`
	RemindDate *time.Time `json:"remind_date"`
}

func (r *RemindUpdateInput) Validate() error {
	if (r.Msg == nil) && (r.RemindDate == nil) && (r.Title == nil) {
		return errors.New("update structure is empty")
	}
	return nil
}
