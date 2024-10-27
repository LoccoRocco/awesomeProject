package service

import "awesomeProject/internal/models"

type repoActor interface {
	CreateActor(actorModel models.CreateActor) (int, error)
	DeleteActor(actorModel models.Actor) error
	UpdateActor(actorModel models.UpdateActor) error
	GetActors(filter models.Actor) ([]models.Actor, error)
}

type Actor struct {
	repo repoActor
}

func NewActor(repo repoActor) *Actor {
	return &Actor{repo}
}

func (a *Actor) CreateActor(actorModel models.CreateActor) (int, error) {
	return a.CreateActor(actorModel)
}

func (a *Actor) DeleteActor(actorModel models.Actor) error {
	return a.DeleteActor(actorModel)
}

func (a *Actor) UpdateActor(actorModel models.UpdateActor) error {
	return a.UpdateActor(actorModel)
}

func (a *Actor) GetActors(filter models.Actor) ([]models.Actor, error) {
	return a.GetActors(filter)
}
