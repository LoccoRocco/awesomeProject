package repository

import "database/sql"

type actor struct {
	db *sql.DB
}

func NewActor(db *sql.DB) *actor {
	return &actor{db}
}
func (a *actor) CreateActor(name string, birthDate string, gender string) (int, error) {
	var actorID int
	err := a.db.QueryRow("INSERT INTO actors (name, birth_date, gender) VALUES($1, $2, $3) RETURNING id", name, birthDate, gender).Scan(&actorID)
	if err != nil {
		return 0, err
	}
	return actorID, nil
}

func (a *actor) UpdateActor(id int, name string, birthDate string, gender string) error {
	_, err := a.db.Exec("UPDATE actors SET name=$1, birth_date=$2, gender=$3 WHERE id=$4", name, birthDate, gender, id)
	return err
}

func (a *actor) DeleteActor(id int) error {
	_, err := a.db.Exec("DELETE FROM actors WHERE id=$1", id)
	return err
}
