package service

import "awesomeProject/internal/models"

type repoMovie interface {
	CreateMovie(movieModel models.CreateMovie) (int, error)
	UpdateMovie(movieModel models.UpdateMovie) error
	DeleteMovie(movieModel models.Movie) error
	GetMovie(name string) (models.Movie, error)
}

type Movie struct {
	repo repoMovie
}

func NewMovie(repo repoMovie) *Movie {
	return &Movie{repo}
}

func (m *Movie) CreateMovie(movieModel models.CreateMovie) (int, error) {
	return m.CreateMovie(movieModel)
}

func (m *Movie) UpdateMovie(movieModel models.UpdateMovie) error {
	return m.UpdateMovie(movieModel)
}

func (m *Movie) DeleteMovie(movieModel models.Movie) error {
	return m.DeleteMovie(movieModel)
}

func (m *Movie) GetMovie(name string) (models.Movie, error) {
	return m.GetMovie(name)
}
