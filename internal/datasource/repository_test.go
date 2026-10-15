package datasource

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"tictactoe/internal/domain"
)

func TestDatasourceRepository(t *testing.T) {
	storage := NewStorage()
	repo := NewGameRepository(storage)

	ctx := context.Background()
	gameID := uuid.New()

	testGame := domain.Game{
		ID: gameID,
		Board: domain.Board{
			{1, 0, 2},
			{0, 1, 0},
			{0, 0, 2},
		},
		Status: domain.StatusInProgress,
	}

	// 1. Get non-existent
	_, err := repo.Get(ctx, gameID)
	if err != domain.ErrGameNotFound {
		t.Fatalf("expected ErrGameNotFound, got %v", err)
	}

	// 2. Save
	if err := repo.Save(ctx, testGame); err != nil {
		t.Fatalf("unexpected save error: %v", err)
	}

	// 3. Get existing
	retrieved, err := repo.Get(ctx, gameID)
	if err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}

	if retrieved.ID != testGame.ID {
		t.Errorf("expected ID %v, got %v", testGame.ID, retrieved.ID)
	}

	if retrieved.Board != testGame.Board {
		t.Errorf("expected board %v, got %v", testGame.Board, retrieved.Board)
	}

	if retrieved.Status != testGame.Status {
		t.Errorf("expected status %v, got %v", testGame.Status, retrieved.Status)
	}
}
