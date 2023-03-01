package repository

import (
	"fmt"
	"github.com/aalmat/todo/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db}
}

func (t *TodoListPostgres) Create(userId int, todoList models.TodoList) (int, error) {
	transaction, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var list_id int
	listQuery := fmt.Sprintf("INSERT INTO %s (description, title) VALUES($1, $2) RETURNING id", todoListsTable)
	row := transaction.QueryRow(listQuery, todoList.Description, todoList.Title)
	if err := row.Scan(&list_id); err != nil {
		transaction.Rollback()
		return 0, err
	}

	userListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES($1, $2)", usersListTable)
	_, err = transaction.Exec(userListQuery, userId, list_id)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	return list_id, transaction.Commit()

}

func (t *TodoListPostgres) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id where ul.user_id=$1", todoListsTable, usersListTable)

	err := t.db.Select(&lists, query, userId)
	//fmt.Println(lists)
	return lists, err

}

func (t *TodoListPostgres) GetListById(userId, listId int) (models.TodoList, error) {
	var list models.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id where ul.user_id=$1 and ul.list_id=$2", todoListsTable, usersListTable)

	err := t.db.Get(&list, query, userId, listId)

	return list, err
}

func (t *TodoListPostgres) UpdateListById(userId, listId int, input models.UpdateList) error {
	ind := 1
	titdes := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Title != nil {
		titdes = append(titdes, fmt.Sprintf("title=$%d", ind))
		args = append(args, *input.Title)
		ind++
	}

	if input.Description != nil {
		titdes = append(titdes, fmt.Sprintf("description=$%d", ind))
		args = append(args, *input.Description)
		ind++
	}

	setValue := strings.Join(titdes, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.list_id and ul.user_id=$%d and ul.list_id=$%d",
		todoListsTable, setValue, usersListTable, ind, ind+1)
	args = append(args, userId, listId)
	_, err := t.db.Exec(query, args...)

	return err
}

func (t *TodoListPostgres) DeleteListById(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", todoListsTable, usersListTable)
	_, err := t.db.Exec(query, userId, listId)

	return err
}
