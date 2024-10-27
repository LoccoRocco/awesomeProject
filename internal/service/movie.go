package service

import "awesomeProject/internal/models"

type repoMovie interface {
	CreateMovie(movieModel models.CreateMovie) (int, error)
	UpdateMovie(movieModel models.UpdateMovie) error
	DeleteMovie(movieModel models.Movie) error
	GetMovie(name string) (models.Movie, error)
}

type movie struct {
	repo repoMovie
}

func NewMovie(repo repoMovie) *movie {
	return &movie{repo}
}

func (m *movie) CreateMovie(movieModel models.CreateMovie) (int, error) {
	return m.CreateMovie(movieModel)
}

func (m *movie) UpdateMovie(movieModel models.UpdateMovie) error {
	return m.UpdateMovie(movieModel)
}

func (m *movie) DeleteMovie(movieModel models.Movie) error {
	return m.DeleteMovie(movieModel)
}

func (m *movie) GetMovie(name string) (models.Movie, error) {
	return m.GetMovie(name)
}
