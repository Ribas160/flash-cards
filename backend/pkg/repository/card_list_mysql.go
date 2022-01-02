package repository

import (
	"fmt"
	"strings"

	app "github.com/Ribas160/flash-cards"
	"github.com/jmoiron/sqlx"
)

type CardListMysql struct {
	db *sqlx.DB
}

func NewCardListMysql(db *sqlx.DB) *CardListMysql {
	return &CardListMysql{db: db}
}

func (r *CardListMysql) Create(userId int, list app.CardList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int64
	createListQiery := fmt.Sprintf("INSERT INTO %s (title) VALUES (?)", cardListsTable)
	row, err := tx.Exec(createListQiery, list.Title)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err = row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListTable := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES(?, ?)", usersListsTable)
	_, err = tx.Exec(createUsersListTable, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}

func (r *CardListMysql) GetAll(userId int) ([]app.CardList, error) {
	var lists []app.CardList
	query := fmt.Sprintf("SELECT cl.id, cl.title FROM %s cl INNER JOIN %s ul ON cl.id = ul.list_id WHERE ul.user_id = ?", cardListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *CardListMysql) GetById(userId int, listId int) (app.CardList, error) {
	var list app.CardList
	query := fmt.Sprintf("SELECT cl.id, cl.title FROM %s cl INNER JOIN %s ul ON cl.id = ul.list_id WHERE ul.user_id = ? AND ul.list_id = ?", cardListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *CardListMysql) Delete(userId int, listId int) error {
	query := fmt.Sprintf("DELETE cl FROM %s cl INNER JOIN %s ul ON cl.id = ul.list_id WHERE cl.id = ul.list_id AND ul.user_id = ? AND ul.list_id = ?", cardListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *CardListMysql) Update(userId int, listId int, input app.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Title != nil {
		setValues = append(setValues, "title = ?")
		args = append(args, *input.Title)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s cl INNER JOIN %s ul ON cl.id = ul.list_id SET %s WHERE cl.id = ul.list_id AND ul.list_id = ? AND ul.user_id = ?", cardListsTable, usersListsTable, setQuery)
	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}
