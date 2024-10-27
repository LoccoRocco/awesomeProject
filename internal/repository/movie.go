package repository

import (
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
)

type movie struct {
	db *sql.DB
}

func NewMovie(db *sql.DB) *movie {
	return &movie{db}
}
func (m *movie) CreateMovie(movieModel models.CreateMovie) (int, error) {
	var movieID int
	err := m.db.QueryRow("INSERT INTO movies(title, release_date, description) VALUES($1, $2, $3) RETURNING id", movieModel.Title, movieModel.ReleaseDate, movieModel.Description).Scan(&movieID)
	if err != nil {
		return 0, err
	}
	return movieID, nil
}

func (m *movie) UpdateMovie(movieModel models.UpdateMovie) error {
	_, err := m.db.Exec("UPDATE movies SET title=$1, release_date=$2, description=$3 WHERE id=$4", movieModel.Title, movieModel.ReleaseDate, movieModel.Description, movieModel.ID)
	return err
}

func (m *movie) DeleteMovie(movieModel models.Movie) error {
	_, err := m.db.Exec("DELETE FROM movies WHERE id=$1", movieModel.ID)
	return err
}

func (m *movie) GetMovie(name string) (models.Movie, error) {
	var movie models.Movie
	err := m.db.QueryRow("SELECT id, title, release_date, description FROM movies WHERE id = $1", name).Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Movie{}, errors.New("movie not found")
		}
		return models.Movie{}, err
	}
	return movie, nil
}
