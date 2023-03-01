package service

import (
	"errors"
	"github.com/aalmat/todo/models"
	"github.com/aalmat/todo/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(r repository.Authorization) *AuthService {
	return &AuthService{r}
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	hash, err := generatePassword(user.PasswordHash)
	if err != nil {
		return 0, err
	}
	user.PasswordHash = hash
	return a.repos.CreateUser(user)
}

const (
	TokenTime = 12 * time.Hour
	signinKey = "ado#123%6fadpK"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	//hash, err := generatePassword(password)
	//if err != nil {
	//	return "", err
	//}
	user, err := a.repos.GetUser(username, password)

	if err != nil {
		return "", err
	}
	//fmt.Println(password)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signinKey))

}

func (a *AuthService) ParseToken(token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign in method")
		}

		return []byte(signinKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := t.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return claims.UserId, nil
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
