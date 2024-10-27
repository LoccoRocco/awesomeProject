package repository

import (
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

type actor struct {
	db *sql.DB
}

func NewActor(db *sql.DB) *actor {
	return &actor{db}
}
func (a *actor) CreateActor(actorModel models.CreateActor) (int, error) {
	var actorID int
	err := a.db.QueryRow("INSERT INTO actors (name, birth_date, gender) VALUES($1, $2, $3)", actorModel.Name, actorModel.BirthDate, actorModel.Gender).Scan(&actorID)
	if err != nil {
		return 0, err
	}
	return actorID, nil
}

func (a *actor) UpdateActor(actorModel models.UpdateActor) error {
	_, err := a.db.Exec("UPDATE actors SET name=$1, birth_date=$2, gender=$3 WHERE id=$4", actorModel.Name, actorModel.BirthDate, actorModel.Gender, actorModel.ID)
	return err
}

func (a *actor) DeleteActor(actorModel models.Actor) error {
	_, err := a.db.Exec("DELETE FROM actors WHERE id=$1", actorModel.ID)
	return err
}

func (a *actor) GetActor(name string) (models.Actor, error) {
	var actor models.Actor
	err := a.db.QueryRow("SELECT id, name, birth_date, gender FROM actors WHERE name = $1", name).Scan(&actor.ID, &actor.Name, &actor.BirthDate, &actor.Gender)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return actor, fmt.Errorf("no actor found with id %d", name)
		}
		return actor, err
	}
	return actor, nil
}
