package models

type User struct {
	Id           int    `json:"-" db:"id"`
	Name         string `json:"name" db:"name" binding:"required"`
	Username     string `json:"username" db:"username" binding:"required"`
	PasswordHash string `json:"password" db:"password_hash" binding:"required"`
}
