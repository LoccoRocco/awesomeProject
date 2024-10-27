package controller

import "awesomeProject/internal/models"

type ServiceMovie interface {
	CreateMovie(movieModel models.CreateMovie) (int, error)
	UpdateMovie(movieModel models.UpdateMovie) error
	DeleteMovie(movieModel models.Movie) error
}
