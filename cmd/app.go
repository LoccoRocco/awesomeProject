package main

import (
	"awesomeProject/internal/postgres"
	"awesomeProject/internal/repository"
	"database/sql"
	"fmt"
	"log"
)

func Run() error {
	db, err := postgres.ConnectToDB()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	actorRepo := repository.NewActor(db)
	movieRepo := repository.NewMovie(db)

	actor, err := actorRepo.CreateActor("Brad", "2000-02-02", "Male")
	if err != nil {
		return err
	}

	movie, err := movieRepo.CreateMovie("TROY", "2000-02-02", "Brad Pit is a troy leader")
	if err != nil {
		return err
	}

	fmt.Printf("actor with id:%v created\n", actor)
	fmt.Printf("movie with id:%v created\n", movie)

	return nil
}
