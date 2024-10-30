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
	mux := http.NewServeMux()
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
	mux.HandleFunc("GET /actors", actorController.GetActors)
	mux.HandleFunc("POST /actors", actorController.CreateActor)
	mux.HandleFunc("PATCH /actors", actorController.UpdateActor)
	mux.HandleFunc("DELETE /actors", actorController.DeleteActor)

	mux.HandleFunc("GET /movies", movieController.GetMovies)
	mux.HandleFunc("POST /movies", movieController.CreateMovie)
	mux.HandleFunc("PATCH /movies", movieController.UpdateMovie)
	mux.HandleFunc("DELETE /movies", movieController.DeleteMovie)

	log.Printf("Сервер запущен на порту %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

	return nil
}
