package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/olenka--91/reminder-app/internal/domain"
)

type Remind interface {
	Create(userID int, rem domain.Remind) (int, error)
	GetByID(userID int, remindID int) (domain.Remind, error)
	GetAll(userID int) ([]domain.Remind, error)
	Delete(userID, remindID int) error
	Update(userID, remindID int, input domain.RemindUpdateInput) error
}

type Repository struct {
	Remind
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Remind: NewRemindPostgres(db)}
}
