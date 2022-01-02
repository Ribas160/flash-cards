package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	cardListsTable  = "lists"
	itemsTable      = "items"
	usersListsTable = "users_lists"
	listItemsTable  = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}

func NewMysqlDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
