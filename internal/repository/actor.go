package repository

import (
	"awesomeProject/internal/models"
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type actor struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
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

func (a *actor) GetActors(filter models.Actor) ([]models.Actor, error) {
	var actors []models.Actor

	query := a.builder.Select("id", "name", "birth_date", "gender").
		From("actors")

	if filter.Name != "" {
		query = query.Where(squirrel.Like{"name": "%" + filter.Name + "%"})
	}
	if filter.BirthDate != "" {
		query = query.Where(squirrel.Eq{"birth_date": filter.BirthDate})
	}
	if filter.Gender != "" {
		query = query.Where(squirrel.Eq{"gender": filter.Gender})
	}

	sqlx, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := a.db.Query(sqlx, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var actor models.Actor
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.BirthDate, &actor.Gender); err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actors, nil
}
