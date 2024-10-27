package controller

type ServiceActor interface {
	CreateActor(name string, birthDate string, gender string) (int, error)
	DeleteActor(id int) error
	UpdateActor(id int, name string, birthDate string, gender string) error
}
