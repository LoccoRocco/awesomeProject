package repository

import (
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
)

type user struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
}

func NewUser(db *sql.DB) *user {
	return &user{db: db, builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}

func (r *user) Register(username string, password []byte) (int64, error) {
	query := r.builder.Insert("users").
		Columns("username", "password", "role").
		Values(username, password, "User").
		RunWith(r.db)

	var userID int64
	err := query.QueryRow().Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *user) Login(username string) (models.User, error) {
	var user models.User

	err := r.builder.Select("*").
		From("users").
		Where("username = ?", username).
		RunWith(r.db).
		QueryRow().Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return user, nil
}

func (r *user) GetUserById(userID int64) (models.User, error) {
	var user models.User

	err := r.builder.Select("username, role").
		From("users").
		Where("id = ?", userID).
		RunWith(r.db).
		QueryRow().Scan(&user.ID, &user.Username, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return models.User{}, err
	}

	return user, nil
}

func (r *user) GetUserByUsername(username string) (models.User, bool, error) {
	var user models.User
	err := r.builder.Select("id, username, role").
		From("users").
		Where("username = ?", username).
		RunWith(r.db).
		QueryRow().Scan(&user.ID, &user.Username, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Возвращаем пустого пользователя и false, если пользователь не найден
			return models.User{}, false, nil
		}
		return models.User{}, false, err
	}

	return user, true, nil
}
