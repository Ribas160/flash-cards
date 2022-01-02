package service

import (
	app "github.com/Ribas160/flash-cards"
	"github.com/Ribas160/flash-cards/pkg/repository"
)

type CardItemService struct {
	repo     repository.CardItem
	listRepo repository.CardList
}

func NewCardItemService(repo repository.CardItem, listRepo repository.CardList) *CardItemService {
	return &CardItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *CardItemService) Create(userId int, listId int, item app.CardItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// List does not existis or does not belong to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *CardItemService) GetAll(userId int, listId int) ([]app.CardItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *CardItemService) GetById(userId int, itemId int) (app.CardItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *CardItemService) Delete(userId int, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *CardItemService) Update(userId int, itemId int, input app.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, itemId, input)
}
