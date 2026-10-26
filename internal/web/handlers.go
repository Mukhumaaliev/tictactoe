package web

import (
	"encoding/json"
	"net/http"
	"tictactoe/internal/domain"
)

type GameHandler struct {
	service domain.GameService
}

func NewGameHandler(service domain.GameService) *GameHandler {
	return &GameHandler{
		service: service,
	}
}

func (h *GameHandler) Create(w http.ResponseWriter, r *http.Request) {
	game, err := h.service.CreateGame(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	wgame := DomainToWeb(game)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(wgame); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *GameHandler) MakeMove(w http.ResponseWriter, r *http.Request) {
	uuidstr := r.PathValue("uuid")
	var wgame *Game
	if err := json.NewDecoder(r.Body).Decode(&wgame); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dgame, err := h.service.NextMove(r.Context(), uuidstr, WebToDomain(wgame))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	wgame = DomainToWeb(dgame)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(wgame); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
