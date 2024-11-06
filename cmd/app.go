package main

import (
	"awesomeProject/internal/controller"
	"awesomeProject/internal/middleware"
	"awesomeProject/internal/postgres"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
	"database/sql"
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
	userRepo := repository.NewUser(db)

	actorService := service.NewActor(actorRepo)
	movieService := service.NewMovie(movieRepo)
	userService := service.NewUser(userRepo)

	actorController := controller.NewActorController(actorService)
	movieController := controller.NewMovieController(movieService)
	userController := controller.NewUserController(userService)

	// Создание HTTP роутера

	mux.HandleFunc("POST /login", userController.Login)
	mux.HandleFunc("POST /register", userController.Register)

	mux.HandleFunc("GET /actors", actorController.GetActors)
	mux.HandleFunc("POST /actors", middleware.Authenticate(actorController.CreateActor))
	mux.HandleFunc("PATCH /actors", middleware.Authenticate(actorController.UpdateActor))
	mux.HandleFunc("DELETE /actors", middleware.Authenticate(actorController.DeleteActor))

	mux.HandleFunc("GET /movies", movieController.GetMovies)
	mux.HandleFunc("POST /movies", middleware.Authenticate(movieController.CreateMovie))
	mux.HandleFunc("PATCH /movies", middleware.Authenticate(movieController.UpdateMovie))
	mux.HandleFunc("DELETE /movies", middleware.Authenticate(movieController.DeleteMovie))

	log.Printf("Сервер запущен на порту %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

	return nil
}
