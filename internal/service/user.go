package service

import (
	"awesomeProject/internal/auth/jwt"
	"awesomeProject/internal/models"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type repoUser interface {
	Register(username string, password []byte) (int64, error)
	Login(username string) (models.User, error)
	GetUserById(userID int64) (models.User, error)
	GetUserByUsername(username string) (models.User, bool, error)
}

type User struct {
	repo repoUser
}

type LoginResponse struct {
	User    models.User
	Access  string
	Refresh string
}

func NewUser(repo repoUser) *User {
	return &User{repo: repo}
}

func (s *User) Register(username, password string) (int64, error) {
	_, exists, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errors.New("user already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	return s.repo.Register(username, passwordHash)
}

func (s *User) Login(username, password string) (LoginResponse, error) {
	user, err := s.repo.Login(username)
	logResponse := LoginResponse{}
	if err != nil {
		return logResponse, err
	}

	fmt.Println(user, "user")

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return logResponse, errors.New("invalid username or password")
	}

	signingKey := []byte("awesomeProject")
	accToken, refToken, err := jwt.GeneratePair(strconv.Itoa(user.ID), signingKey)

	logResponse.Access = accToken
	logResponse.Refresh = refToken
	logResponse.User = user

	return logResponse, nil
}

func (s *User) GetUserById(userID int64) (models.User, error) {
	return s.repo.GetUserById(userID)
}
