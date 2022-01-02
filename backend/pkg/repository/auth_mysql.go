package repository

import (
	"fmt"

	app "github.com/Ribas160/flash-cards"
	"github.com/jmoiron/sqlx"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user app.User) (int, error) {
	var id int64

	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) VALUES (?, ?, ?)", usersTable)

	row, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err = row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthMysql) GetUser(username string, password string) (app.User, error) {
	var user app.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=? AND password_hash=?", usersTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
