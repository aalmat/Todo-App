package repository

import (
	"github.com/aalmat/todo/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetListById(userId, listId int) (models.TodoList, error)
	UpdateListById(userId, listId int, input models.UpdateList) error
	DeleteListById(userId, listId int) error
}

type TodoItem interface {
	CreateItem(userId, listId int, list models.TodoItem) (int, error)
	GetAllItem(userId, listId int) ([]models.TodoItem, error)
	GetItemById(userId, itemId int) (models.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input models.UpdateItem) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
