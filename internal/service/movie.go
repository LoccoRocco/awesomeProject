package service

type storeMovie interface {
	CreateMovie(title string, releaseDate string, description string) (int, error)
	UpdateMovie(id int, title string, releaseDate string, description string) error
	DeleteMovie(id int) error
}

type movie struct {
	store storeMovie
}

func NewMovie(store storeMovie) *movie {
	return &movie{store}
}

func (m *movie) CreateMovie(title string, releaseDate string, description string) (int, error) {
	return m.CreateMovie(title, releaseDate, description)
}

func (m *movie) UpdateMovie(id int, title string, releaseDate string, description string) error {
	return m.UpdateMovie(id, title, releaseDate, description)
}

func (m *movie) DeleteMovie(id int) error {
	return m.DeleteMovie(id)
}
