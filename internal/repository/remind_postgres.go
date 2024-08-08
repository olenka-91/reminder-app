package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/olenka--91/reminder-app/internal/domain"
)

type RemindPostgres struct {
	db *sqlx.DB
}

func NewRemindPostgres(db *sqlx.DB) *RemindPostgres {
	return &RemindPostgres{db: db}
}

func (r *RemindPostgres) Create(userID int, rem domain.Remind) (int, error) {
	return 0, nil
}
func (r *RemindPostgres) GetByID(userID int, remindID int) (domain.Remind, error) {
	return domain.Remind{}, nil
}

func (r *RemindPostgres) GetAll(userID int) ([]domain.Remind, error) {
	return nil, nil
}
func (r *RemindPostgres) Delete(userID, remindID int) error {
	return nil
}
func (r *RemindPostgres) Update(userID int, input domain.RemindUpdateInput) error {
	return nil
}
