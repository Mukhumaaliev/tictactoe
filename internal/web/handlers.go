package web

import (
	"encoding/json"
	"net/http"
	"tictactoe/internal/domain"
	"uuid"
)

type GameHandler struct {
	service domain.GameService
	repo    domain.GameRepository
}

func NewGameHandler(service domain.GameService, repo domain.GameRepository) *GameHandler {
	return &GameHandler{
		service: service,
		repo:    repo,
	}
}

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	game := domain.NewGame()
	h.repo.Save(game)
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
	uuid, err := uuid.Parse(uuidstr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}
	dgame, err := h.repo.Get(uuid)
	if err != nil {
		http.Error(w, "Не смогли найти игру с таким UUID", http.StatusNotFound)
		return
	}
	if dgame.Status != domain.GameStatusInProgress {
		http.Error(w, "Игра уже завершена", http.StatusBadRequest)
		return
	}

	var wgame Game
	if err := json.NewDecoder(r.Body).Decode(&wgame); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if valid, err := h.service.IsValidBoard(dgame.Board, wgame.Board); !valid {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dgame.Board = wgame.Board
	domain.UpdateGameStatus(dgame)
	if dgame.Status == domain.GameStatusInProgress {
		row, col := h.service.FindBestMove(dgame.Board)
		dgame.Board[row][col] = domain.CellO
		domain.UpdateGameStatus(dgame)
	}
	h.repo.Save(dgame)
	wgame = *DomainToWeb(dgame)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(wgame); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
