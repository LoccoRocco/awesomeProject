package repository

import (
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
)

type actor struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
}

type ActorFilter struct {
	IDIn          []int
	Name          string
	BirthDateFrom string
	BirthDateTo   string
}

func NewActor(db *sql.DB) *actor {
	return &actor{db: db, builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}
func (a *actor) CreateActor(actorModel models.CreateActor) (int, error) {
	var actorID int
	query, args, err := a.builder.
		Insert("actors").
		Columns("name", "birth_date", "gender").
		Values(actorModel.Name, actorModel.BirthDate, actorModel.Gender).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	err = a.db.QueryRow(query, args...).Scan(&actorID)
	if err != nil {
		return 0, err
	}
	return actorID, nil
}

func (a *actor) UpdateActor(actorModel models.UpdateActor) error {
	query, args, err := a.builder.
		Update("actors").
		Set("name", actorModel.Name).
		Set("birth_date", actorModel.BirthDate).
		Set("gender", actorModel.Gender).
		Where("id = ?", actorModel.ID).
		ToSql()

	if err != nil {
		return err
	}

	_, err = a.db.Exec(query, args...)
	return err
}

func (a *actor) DeleteActor(actorModel models.Actor) error {
	query, args, err := a.builder.
		Delete("actors").
		Where("id = ?", actorModel.ID).
		ToSql()

	if err != nil {
		return err
	}

	_, err = a.db.Exec(query, args...)
	return err
}

func (a *actor) GetActors(filter ActorFilter) ([]models.Actor, error) {
	var actors []models.Actor

	query := a.builder.Select("a.id", "a.name", "a.birth_date").
		From("actors a")

	if len(filter.IDIn) > 0 {
		query = query.Where(squirrel.Eq{"a.id": filter.IDIn})
	}

	if filter.Name != "" {
		query = query.Where(squirrel.Like{"a.name": "%" + filter.Name + "%"})
	}

	if filter.BirthDateFrom != "" {
		query = query.Where("a.birth_date >= ?", filter.BirthDateFrom)
	}
	if filter.BirthDateTo != "" {
		query = query.Where("a.birth_date <= ?", filter.BirthDateTo)
	}

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, errors.New("error building SQL query: " + err.Error())
	}

	rows, err := a.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, errors.New("error executing query: " + err.Error())
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var actor models.Actor
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.BirthDate); err != nil {
			return nil, errors.New("error scanning row: " + err.Error())
		}
		actors = append(actors, actor)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("error iterating rows: " + err.Error())
	}

	return actors, nil
}
