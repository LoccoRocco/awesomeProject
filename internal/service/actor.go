package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
)

type repoActor interface {
	CreateActor(actorModel models.CreateActor) (int, error)
	DeleteActor(actorModel models.Actor) error
	UpdateActor(actorModel models.UpdateActor) error
	GetActors(filter repository.ActorFilter) ([]models.Actor, error)
}

type Actor struct {
	repo repoActor
}

func NewActor(repo repoActor) *Actor {
	return &Actor{repo}
}

func (a *Actor) CreateActor(actorModel models.CreateActor) (int, error) {
	return a.repo.CreateActor(actorModel)
}

func (a *Actor) DeleteActor(actorModel models.Actor) error {
	return a.repo.DeleteActor(actorModel)
}

func (a *Actor) UpdateActor(actorModel models.UpdateActor) error {
	return a.repo.UpdateActor(actorModel)
}

func (a *Actor) GetActors(filter repository.ActorFilter) ([]models.Actor, error) {
	return a.repo.GetActors(filter)
}
