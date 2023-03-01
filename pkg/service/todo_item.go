package service

import (
	"github.com/aalmat/todo/models"
	"github.com/aalmat/todo/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo}
}

func (r *TodoItemService) CreateItem(userId, listId int, item models.TodoItem) (int, error) {
	return r.repo.CreateItem(userId, listId, item)
}

func (r *TodoItemService) GetAllItem(userId, listId int) ([]models.TodoItem, error) {
	return r.repo.GetAllItem(userId, listId)
}

func (r *TodoItemService) GetItemById(userId, itemId int) (models.TodoItem, error) {
	return r.repo.GetItemById(userId, itemId)
}

func (r *TodoItemService) DeleteItem(userId, itemId int) error {
	return r.repo.DeleteItem(userId, itemId)
}

func (r *TodoItemService) UpdateItem(userId, itemId int, input models.UpdateItem) error {
	return r.repo.UpdateItem(userId, itemId, input)
}
