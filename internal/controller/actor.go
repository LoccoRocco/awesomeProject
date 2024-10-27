package controller

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type ActorController struct {
	service *service.Actor
}

func NewActorController(service *service.Actor) *ActorController {
	return &ActorController{service: service}
}

func (a *ActorController) GetActors(w http.ResponseWriter, r *http.Request) {
	var filter repository.ActorFilter

	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actors, err := a.service.GetActors(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(actors); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *ActorController) CreateActor(w http.ResponseWriter, r *http.Request) {
	var actor models.CreateActor

	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := a.service.CreateActor(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *ActorController) UpdateActor(w http.ResponseWriter, r *http.Request) {
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

	if err := a.service.UpdateActor(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *ActorController) DeleteActor(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	var actor models.Actor
	actor.ID = id

	if err := a.service.DeleteActor(actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
