package service

import (
	app "github.com/Ribas160/flash-cards"
	"github.com/Ribas160/flash-cards/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type CardList interface {
	Create(userId int, list app.CardList) (int, error)
	GetAll(userId int) ([]app.CardList, error)
	GetById(userId int, listId int) (app.CardList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input app.UpdateListInput) error
}

type CardItem interface {
	Create(userId int, listId int, item app.CardItem) (int, error)
	GetAll(userId int, listId int) ([]app.CardItem, error)
	GetById(userId int, itemId int) (app.CardItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input app.UpdateItemInput) error
}

type Service struct {
	Authorization
	CardList
	CardItem
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CardList:      NewCardListService(repos.CardList),
		CardItem:      NewCardItemService(repos.CardItem, repos.CardList),
	}
}
