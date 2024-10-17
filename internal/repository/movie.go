package repository

import "database/sql"

type movie struct {
	db *sql.DB
}

func NewMovie(db *sql.DB) *movie {
	return &movie{db}
}

func (m *movie) CreateMovie(title string, releaseDate string, description string) (int, error) {
	var movieID int
	err := m.db.QueryRow("INSERT INTO movies(title, release_date, description) VALUES($1, $2, $3) RETURNING id", title, releaseDate, description).Scan(&movieID)
	if err != nil {
		return 0, err
	}
	return movieID, nil
}

func (m *movie) UpdateMovie(id int, title string, releaseDate string, description string) error {
	_, err := m.db.Exec("UPDATE movies SET title=$1, release_date=$2, description=$3 WHERE id=$4", title, releaseDate, description, id)
	return err
}

func (m *movie) DeleteMovie(id int) error {
	_, err := m.db.Exec("DELETE FROM movies WHERE id=$1", id)
	return err
}
