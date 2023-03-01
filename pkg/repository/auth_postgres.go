package repository

import (
	//"errors"
	"fmt"
	"github.com/aalmat/todo/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (a *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	//fmt.Println(user.PasswordHash)
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := a.db.QueryRow(query, user.Name, user.Username, user.PasswordHash)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 LIMIT 1", usersTable)
	err := a.db.Get(&user, query, username)
	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
