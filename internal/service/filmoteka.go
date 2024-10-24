package controller

type filmoteka struct {
	movie serviceMovie
	actor serviceActor
}

func NewFilmoteka(movie serviceMovie, actor serviceActor) *filmoteka {
	return &filmoteka{movie: movie, actor: actor}
}
