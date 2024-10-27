package service

type storeActor interface {
	CreateActor(name string, birthDate string, gender string) (int, error)
	DeleteActor(id int) error
	UpdateActor(id int, name string, birthDate string, gender string) error
}

type actor struct {
	store storeActor
}

func NewActor(store storeActor) *actor {
	return &actor{store}
}

func (a *actor) CreateActor(name string, birthDate string, gender string) (int, error) {
	return a.CreateActor(name, birthDate, gender)
}

func (a *actor) DeleteActor(id int) error {
	return a.DeleteActor(id)
}

func (a *actor) UpdateActor(id int, name string, birthDate string, gender string) error {
	return a.UpdateActor(id, name, birthDate, gender)
}
