package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/sirupsen/logrus"
)

type RemindPostgres struct {
	db *sqlx.DB
}

func NewRemindPostgres(db *sqlx.DB) *RemindPostgres {
	return &RemindPostgres{db: db}
}

func (r *RemindPostgres) Create(userID int, rem domain.Remind) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	queryString := fmt.Sprintf("INSERT INTO %s (title, msg, remind_date) VALUES ($1,$2,$3) RETURNING id", remindTable)
	row := tx.QueryRow(queryString, rem.Title, rem.Msg, rem.RemindDate)
	var id int
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
func (r *RemindPostgres) GetByID(userID int, remindID int) (domain.Remind, error) {
	var rem domain.Remind
	queryString := fmt.Sprintf("SELECT t.id, t.title, t.msg, t.remind_date as RemindDate FROM %s t WHERE t.id=$1", remindTable)
	err := r.db.Get(&rem, queryString, remindID)

	return rem, err
}

func (r *RemindPostgres) GetAll(userID int) ([]domain.Remind, error) {
	var rem []domain.Remind
	queryString := fmt.Sprintf("SELECT t.id, t.title, t.msg, t.remind_date as RemindDate FROM %s t", remindTable)
	err := r.db.Select(&rem, queryString)

	return rem, err
}

func (r *RemindPostgres) Delete(userID, remindID int) error {
	queryString := fmt.Sprintf("DELETE FROM %s t WHERE t.id=$1", remindTable)
	_, err := r.db.Exec(queryString, remindID)
	return err
}

func (r *RemindPostgres) Update(userID, remindID int, input domain.RemindUpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argIDs := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argIDs))
		args = append(args, *input.Title)
		argIDs++
	}

	if input.Msg != nil {
		setValues = append(setValues, fmt.Sprintf("msg=$%d", argIDs))
		args = append(args, *input.Msg)
		argIDs++
	}

	if input.RemindDate != nil {
		setValues = append(setValues, fmt.Sprintf("remind_date=$%d", argIDs))
		args = append(args, *input.RemindDate)
		argIDs++
	}

	updateString := strings.Join(setValues, " ,")
	queryString := fmt.Sprintf("UPDATE %s t SET %s WHERE id = $%d", remindTable, updateString, argIDs)
	args = append(args, remindID)

	logrus.Debugf("updateQuery: %s", queryString)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(queryString, args...)

	return err
}
