package repository

import (
	"fmt"
	"strings"

	app "github.com/Ribas160/flash-cards"
	"github.com/jmoiron/sqlx"
)

type CardItemMysql struct {
	db *sqlx.DB
}

func NewCardItemMysql(db *sqlx.DB) *CardItemMysql {
	return &CardItemMysql{
		db: db,
	}
}

func (r *CardItemMysql) Create(listId int, item app.CardItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int64
	createItemQuery := fmt.Sprintf("INSERT INTO %s (lang1, lang2) VALUES (?, ?)", itemsTable)
	row, err := tx.Exec(createItemQuery, item.Lang1, item.Lang2)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	itemId, err = row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES (?, ?)", listItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(itemId), tx.Commit()
}

func (r *CardItemMysql) GetAll(userId int, listId int) ([]app.CardItem, error) {
	var items []app.CardItem

	query := fmt.Sprintf("SELECT ci.id, ci.lang1, ci.lang2 FROM %s ci INNER JOIN %s li ON li.item_id = ci.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE li.list_id = ? AND ul.user_id = ?",
		itemsTable, listItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *CardItemMysql) GetById(userId int, itemId int) (app.CardItem, error) {
	var item app.CardItem

	query := fmt.Sprintf("SELECT ci.id, ci.lang1, ci.lang2 FROM %s ci INNER JOIN %s li ON li.item_id = ci.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ci.id = ? AND ul.user_id = ?",
		itemsTable, listItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *CardItemMysql) Delete(userId int, itemId int) error {
	query := fmt.Sprintf("DELETE ci FROM %s ci INNER JOIN %s li ON li.item_id = ci.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ci.id = ? AND ul.user_id = ?",
		itemsTable, listItemsTable, usersListsTable)
	_, err := r.db.Exec(query, itemId, userId)

	return err
}

func (r *CardItemMysql) Update(userId int, itemId int, input app.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Lang1 != nil {
		setValues = append(setValues, "lang1 = ?")
		args = append(args, *input.Lang1)
	}

	if input.Lang2 != nil {
		setValues = append(setValues, "lang2 = ?")
		args = append(args, *input.Lang2)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ci INNER JOIN %s li ON li.item_id = ci.id INNER JOIN %s ul ON ul.list_id = li.list_id SET %s 
						WHERE ci.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = ? AND ci.id = ?`,
		itemsTable, listItemsTable, usersListsTable, setQuery)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)

	return err
}
