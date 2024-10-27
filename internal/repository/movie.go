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

type MovieFilter struct {
	IDIn            []int
	Title           string
	ReleaseDateFrom string
	ReleaseDateTo   string
	ActorIDIn       []int
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

func (m *movie) GetMovies(filter MovieFilter) ([]models.Movie, error) {
	var movies []models.Movie

	query := m.builder.Select("m.id", "m.title", "m.release_date").
		From("movies m")

	if len(filter.IDIn) > 0 {
		query = query.Where(squirrel.Eq{"m.id": filter.IDIn})
	}

	if filter.Title != "" {
		query = query.Where(squirrel.Like{"m.title": "%" + filter.Title + "%"})
	}

	if filter.ReleaseDateFrom != "" {
		query = query.Where("m.release_date >= ?", filter.ReleaseDateFrom)
	}
	if filter.ReleaseDateTo != "" {
		query = query.Where("m.release_date <= ?", filter.ReleaseDateTo)
	}

	if len(filter.ActorIDIn) > 0 {
		subQuery := squirrel.Select("ma.movie_id").
			From("movie_actors ma").
			Where(squirrel.Eq{"ma.actor_id": filter.ActorIDIn})

		query = query.Where(squirrel.Eq{"m.id": subQuery})
	}

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := m.db.Query(sqlQuery, args...)
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
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
