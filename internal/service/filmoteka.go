package service

import "awesomeProject/internal/controller"

type filmoteka struct {
	movie controller.ServiceMovie
	actor controller.ServiceActor
}

func NewFilmoteka(movie controller.ServiceMovie, actor controller.ServiceActor) *filmoteka {
	return &filmoteka{movie: movie, actor: actor}
}
