package controller

import "awesomeProject/internal/models"

type ServiceActor interface {
	CreateActor(actorModel models.CreateActor) (int, error)
	DeleteActor(actorModel models.Actor) error
	UpdateActor(actorModel models.UpdateActor) error
	GetActor(name string) (models.Actor, error)
}
