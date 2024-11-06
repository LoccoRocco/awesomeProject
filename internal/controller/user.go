package controller

import (
	"awesomeProject/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserController struct {
	service *service.User
}

func NewUserController(service *service.User) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, err := c.service.Register(userInput.Username, userInput.Password)
	if err != nil {
		if err.Error() == "user already exists" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]int64{"user_id": userID})
	if err != nil {
		http.Error(w, "Internal server error Register encode", http.StatusInternalServerError)
		return
	}
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	loginResponse, err := c.service.Login(userInput.Username, userInput.Password)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "invalid credentials" {
			status = http.StatusUnauthorized
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(loginResponse); err != nil {
		http.Error(w, "Internal server error Login", http.StatusInternalServerError)
		return
	}
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := c.service.GetUserById(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Can't Encode", http.StatusBadRequest)
		return
	}
}
