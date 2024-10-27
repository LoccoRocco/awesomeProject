package controller

type ServiceMovie interface {
	CreateMovie(title string, releaseDate string, description string) (int, error)
	UpdateMovie(id int, title string, releaseDate string, description string) error
	DeleteMovie(id int) error
}
