package controller

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type Filmoteka struct {
	movie *service.Movie
	actor *service.Actor
}

func NewFilmoteka(movie *service.Movie, actor *service.Actor) *Filmoteka {
	return &Filmoteka{movie: movie, actor: actor}
}

func (f *Filmoteka) GetActors(w http.ResponseWriter, r *http.Request) {
	var filter repository.ActorFilter

	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actors, err := f.actor.GetActors(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(actors); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (f *Filmoteka) GetMovies(w http.ResponseWriter, r *http.Request) {
	var filter repository.MovieFilter

	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movies, err := f.movie.GetMovies(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (f *Filmoteka) CreateActor(w http.ResponseWriter, r *http.Request) {
	var actor models.CreateActor

	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := f.actor.CreateActor(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (f *Filmoteka) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.CreateMovie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := f.movie.CreateMovie(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (f *Filmoteka) UpdateActor(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	var actor models.UpdateActor
	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actor.ID = id

	if err := f.actor.UpdateActor(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (f *Filmoteka) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var movie models.UpdateMovie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movie.ID = id

	if err := f.movie.UpdateMovie(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (f *Filmoteka) DeleteActor(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	var actor models.Actor
	actor.ID = id

	if err := f.actor.DeleteActor(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (f *Filmoteka) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var movie models.Movie
	movie.ID = id

	if err := f.movie.DeleteMovie(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
