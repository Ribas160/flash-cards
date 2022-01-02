package repository

import (
	app "github.com/Ribas160/flash-cards"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(username string, password string) (app.User, error)
}

type CardList interface {
	Create(userId int, list app.CardList) (int, error)
	GetAll(userId int) ([]app.CardList, error)
	GetById(userId int, listId int) (app.CardList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input app.UpdateListInput) error
}

type CardItem interface {
	Create(listId int, item app.CardItem) (int, error)
	GetAll(userId int, listId int) ([]app.CardItem, error)
	GetById(userId int, itemId int) (app.CardItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input app.UpdateItemInput) error
}

type Repository struct {
	Authorization
	CardList
	CardItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		CardList:      NewCardListMysql(db),
		CardItem:      NewCardItemMysql(db),
	}
}
