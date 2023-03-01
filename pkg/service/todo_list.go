package service

import (
	"github.com/aalmat/todo/models"
	"github.com/aalmat/todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo}
}

func (r *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return r.repo.Create(userId, list)
}

func (r *TodoListService) GetAll(userId int) ([]models.TodoList, error) {
	return r.repo.GetAll(userId)
}

func (r *TodoListService) GetListById(userId, listId int) (models.TodoList, error) {
	return r.repo.GetListById(userId, listId)
}

func (r *TodoListService) UpdateListById(userId, listId int, input models.UpdateList) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return r.repo.UpdateListById(userId, listId, input)
}

func (r *TodoListService) DeleteListById(userId, listId int) error {
	return r.repo.DeleteListById(userId, listId)
}
