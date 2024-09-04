package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/olenka--91/reminder-app/internal/domain"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(u domain.User) (int, error) {
	var id int
	queryStr := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := a.db.QueryRow(queryStr, u.Name, u.Username, u.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (domain.User, error) {
	var user domain.User
	queryStr := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", userTable)
	err := a.db.Get(&user, queryStr, username, password)

	return user, err
}
