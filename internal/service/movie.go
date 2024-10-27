package service

import "awesomeProject/internal/models"

type storeMovie interface {
	CreateMovie(movieModel models.CreateMovie) (int, error)
	UpdateMovie(movieModel models.UpdateMovie) error
	DeleteMovie(movieModel models.Movie) error
}

type movie struct {
	store storeMovie
}

func NewMovie(store storeMovie) *movie {
	return &movie{store}
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
