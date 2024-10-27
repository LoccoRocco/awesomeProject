package controller

import "awesomeProject/internal/service"

type filmoteka struct {
	movie *service.Movie
	actor *service.Actor
}

func NewFilmoteka(movie *service.Movie, actor *service.Actor) *filmoteka {
	return &filmoteka{movie: movie, actor: actor}
}
