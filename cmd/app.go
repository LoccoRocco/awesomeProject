package main

import (
	"awesomeProject/internal/postgres"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
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

	actorService := service.NewActor(actorRepo)
	movieService := service.NewMovie(movieRepo)

	filmotekaService := service.NewFilmoteka(movieService, actorService)

	fmt.Println(filmotekaService)

	return nil
}
