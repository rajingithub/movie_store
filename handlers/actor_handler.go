package handlers

import (
	"encoding/json"
	"movie_store/services"
	"net/http"
)

type ActorHandler struct {
	service *services.ActorService
}

func NewActorHandler(service *services.ActorService) *ActorHandler {
	return &ActorHandler{service: service}
}

func (h *ActorHandler) GetAllActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.service.GetAllActors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actors); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
