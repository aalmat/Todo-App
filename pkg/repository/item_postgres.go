package repository

import (
	"fmt"
	"github.com/aalmat/todo/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db}
}

func (t *TodoItemPostgres) CreateItem(userId, listId int, item models.TodoItem) (int, error) {
	transaction, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	listQuery := fmt.Sprintf("INSERT INTO %s (title, description, done) VALUES($1, $2, $3) RETURNING id", todoItemsTable)
	row := transaction.QueryRow(listQuery, item.Title, item.Description, item.Done)
	var itemId int
	if err := row.Scan(&itemId); err != nil {
		transaction.Rollback()
		return 0, err
	}

	userListQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES($1, $2)", listsItemsTable)
	_, err = transaction.Exec(userListQuery, listId, itemId)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	return itemId, transaction.Commit()
}

func (t *TodoItemPostgres) GetAllItem(userId, listId int) ([]models.TodoItem, error) {
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done 
								FROM %s ti INNER JOIN %s tl on ti.id=tl.item_id 
								INNER JOIN %s ul on ul.list_id=tl.list_id
								WHERE tl.list_id=$1 and ul.user_id=$2`, todoItemsTable, listsItemsTable, usersListTable)
	var items []models.TodoItem

	err := t.db.Select(&items, query, listId, userId)

	return items, err
}

func (t *TodoItemPostgres) GetItemById(userId, itemId int) (models.TodoItem, error) {
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done 
								FROM %s ti INNER JOIN %s tl ON ti.id=tl.item_id 
							    INNER JOIN %s ul on tl.list_id=ul.list_id 
								WHERE tl.item_id=$1 and ul.user_id=$2`, todoItemsTable, listsItemsTable, usersListTable)

	var item models.TodoItem
	err := t.db.Get(&item, query, itemId, userId)

	return item, err
}

func (t *TodoItemPostgres) DeleteItem(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
								WHERE ti.id=li.item_id AND li.list_id=ul.list_id AND ul.user_id=$1 AND ti.id=$2`,
		todoItemsTable, listsItemsTable, usersListTable)

	_, err := t.db.Exec(query, userId, itemId)
	return err
}

func (t *TodoItemPostgres) UpdateItem(userId, itemId int, input models.UpdateItem) error {
	ind := 1
	titdesdon := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Title != nil {
		titdesdon = append(titdesdon, fmt.Sprintf("title=$%d", ind))
		args = append(args, *input.Title)
		ind++
	}

	if input.Description != nil {
		titdesdon = append(titdesdon, fmt.Sprintf("description=$%d", ind))
		args = append(args, *input.Description)
		ind++
	}

	if input.Done != nil {
		titdesdon = append(titdesdon, fmt.Sprintf("done=$%d", ind))
		args = append(args, *input.Done)
		ind++
	}

	setValue := strings.Join(titdesdon, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul 
                    						WHERE ti.id=li.item_id AND li.list_id=ul.list_id and ul.user_id=$%d AND li.item_id=$%d`,
		todoItemsTable, setValue, listsItemsTable, usersListTable, ind, ind+1)

	args = append(args, userId, itemId)
	_, err := t.db.Exec(query, args...)

	return err
}
