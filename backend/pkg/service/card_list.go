package service

import (
	app "github.com/Ribas160/flash-cards"
	"github.com/Ribas160/flash-cards/pkg/repository"
)

type CardListService struct {
	repo repository.CardList
}

func NewCardListService(repo repository.CardList) *CardListService {
	return &CardListService{repo: repo}
}

func (s *CardListService) Create(userId int, list app.CardList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *CardListService) GetAll(userId int) ([]app.CardList, error) {
	return s.repo.GetAll(userId)
}

func (s *CardListService) GetById(userId int, listId int) (app.CardList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *CardListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *CardListService) Update(userId int, listId int, input app.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
