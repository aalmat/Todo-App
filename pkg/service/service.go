package service

import (
	"github.com/aalmat/todo/models"
	"github.com/aalmat/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetListById(userId, listId int) (models.TodoList, error)
	UpdateListById(userId, listId int, input models.UpdateList) error
	DeleteListById(userId, listId int) error
}

type TodoItem interface {
	CreateItem(userId, listId int, item models.TodoItem) (int, error)
	GetAllItem(userId, listId int) ([]models.TodoItem, error)
	GetItemById(userId, itemId int) (models.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input models.UpdateItem) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem),
	}
}
