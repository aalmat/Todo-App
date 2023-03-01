package models

import (
	"errors"
)

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateList struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (u UpdateList) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("Nothing to change")
	}
	return nil
}

type UpdateItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (u UpdateItem) Validate() error {
	if u.Title == nil && u.Description == nil && u.Done == nil {
		return errors.New("Nothing to change")
	}
	return nil
}
