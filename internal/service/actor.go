package service

import "awesomeProject/internal/models"

type repoActor interface {
	CreateActor(actorModel models.CreateActor) (int, error)
	DeleteActor(actorModel models.Actor) error
	UpdateActor(actorModel models.UpdateActor) error
	GetActor(name string) (models.Actor, error)
}

type actor struct {
	repo repoActor
}

func NewActor(repo repoActor) *actor {
	return &actor{repo}
}

func (a *actor) CreateActor(actorModel models.CreateActor) (int, error) {
	return a.CreateActor(actorModel)
}

func (a *actor) DeleteActor(actorModel models.Actor) error {
	return a.DeleteActor(actorModel)
}

func (a *actor) UpdateActor(actorModel models.UpdateActor) error {
	return a.UpdateActor(actorModel)
}

func (a *actor) GetActor(name string) (models.Actor, error) {
	return a.GetActor(name)
}
