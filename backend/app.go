package app

import "errors"

type CardList struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type CardItem struct {
	Id    int    `json:"id" db:"id"`
	Lang1 string `json:"lang1" db:"lang1" binding:"required"`
	Lang2 string `json:"lang2" db:"lang2" binding:"required"`
}

type UpdateListInput struct {
	Title *string `json:"title"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil {
		return errors.New("Update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Lang1 *string `json:"lang1"`
	Lang2 *string `json:"lang2"`
}

func (i UpdateItemInput) Validate() error {
	if i.Lang1 == nil && i.Lang2 == nil {
		return errors.New("Update structure has no values")
	}

	return nil
}
