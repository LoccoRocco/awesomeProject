package repository

import (
	"awesomeProject/internal/models"
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type movie struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
}

func NewMovie(db *sql.DB) *movie {
	return &movie{db: db, builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}
func (m *movie) CreateMovie(movieModel models.CreateMovie) (int, error) {
	var movieID int

	query, args, err := m.builder.Insert("movies").
		Columns("title", "release_date", "description").
		Values(movieModel.Title, movieModel.ReleaseDate, movieModel.Description).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	err = m.db.QueryRow(query, args...).Scan(&movieID)
	if err != nil {
		return 0, err
	}
	return movieID, nil
}

func (m *movie) UpdateMovie(movieModel models.UpdateMovie) error {
	query, args, err := m.builder.Update("movies").
		Set("title", movieModel.Title).
		Set("release_date", movieModel.ReleaseDate).
		Set("description", movieModel.Description).
		Where(squirrel.Eq{"id": movieModel.ID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = m.db.Exec(query, args...)
	return err
}

func (m *movie) DeleteMovie(movieModel models.Movie) error {
	query, args, err := m.builder.Delete("movies").
		Where(squirrel.Eq{"id": movieModel.ID}).
		ToSql()

	if err != nil {
		return err
	}
	_, err = m.db.Exec(query, args...)
	return err
}

func (m *movie) GetMovies(filter models.Movie) ([]models.Movie, error) {
	var movies []models.Movie

	query := m.builder.Select("id", "title", "release_date", "description").
		From("movies")

	if filter.ID != 0 {
		query = query.Where(squirrel.Eq{"id": filter.ID})
	}
	if filter.Title != "" {
		query = query.Where(squirrel.Like{"title": "%" + filter.Title + "%"})
	}
	if filter.ReleaseDate != 0 {
		query = query.Where(squirrel.Eq{"release_date": filter.ReleaseDate})
	}
	if filter.Description != "" {
		query = query.Where(squirrel.Like{"description": "%" + filter.Description + "%"})
	}

	sqlx, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := m.db.Query(sqlx, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Description); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
