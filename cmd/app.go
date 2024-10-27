package main

import (
	"awesomeProject/internal/controller"
	"awesomeProject/internal/postgres"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var (
	port = ":8080"
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

	actorController := controller.NewActorController(actorService)
	movieController := controller.NewMovieController(movieService)

	filmotekaService := controller.NewFilmoteka(movieService, actorService)

	fmt.Println(filmotekaService)

	// Создание HTTP роутера
	http.HandleFunc("/movies", movieController.GetMovies)
	http.HandleFunc("/movies/create", movieController.CreateMovie)
	http.HandleFunc("/movies/update/", movieController.UpdateMovie)
	http.HandleFunc("/movies/delete/", movieController.DeleteMovie)

	http.HandleFunc("/actors", actorController.GetActors)
	http.HandleFunc("/actors/create", actorController.CreateActor)
	http.HandleFunc("/actors/update/", actorController.UpdateActor)
	http.HandleFunc("/actors/delete/", actorController.DeleteActor)

	log.Printf("Сервер запущен на порту %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

	return nil
}
