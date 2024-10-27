package service

import "awesomeProject/internal/models"

type storeActor interface {
	CreateActor(actorModel models.CreateActor) (int, error)
	DeleteActor(actorModel models.Actor) error
	UpdateActor(actorModel models.UpdateActor) error
}

type actor struct {
	store storeActor
}

func NewActor(store storeActor) *actor {
	return &actor{store}
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
