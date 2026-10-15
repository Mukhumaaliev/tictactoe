package service

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"tictactoe/internal/domain"
)

type mockRepo struct {
	game domain.Game
	err  error
}

func (m *mockRepo) Save(ctx context.Context, game domain.Game) error {
	m.game = game
	return m.err
}

func (m *mockRepo) Get(ctx context.Context, id uuid.UUID) (domain.Game, error) {
	if m.err != nil {
		return domain.Game{}, m.err
	}
	return m.game, nil
}

func TestCheckGameOver(t *testing.T) {
	tests := []struct {
		name       string
		board      domain.Board
		wantStatus domain.GameStatus
		wantOver   bool
	}{
		{
			name: "Row win for X",
			board: domain.Board{
				{1, 1, 1},
				{0, 2, 0},
				{2, 0, 0},
			},
			wantStatus: domain.StatusPlayerWon,
			wantOver:   true,
		},
		{
			name: "Col win for O",
			board: domain.Board{
				{2, 1, 0},
				{2, 1, 0},
				{2, 0, 1},
			},
			wantStatus: domain.StatusAIWon,
			wantOver:   true,
		},
		{
			name: "Diagonal win for X",
			board: domain.Board{
				{1, 2, 0},
				{0, 1, 2},
				{0, 0, 1},
			},
			wantStatus: domain.StatusPlayerWon,
			wantOver:   true,
		},
		{
			name: "Draw",
			board: domain.Board{
				{1, 2, 1},
				{1, 2, 2},
				{2, 1, 1},
			},
			wantStatus: domain.StatusDraw,
			wantOver:   true,
		},
		{
			name: "In progress",
			board: domain.Board{
				{1, 0, 0},
				{0, 2, 0},
				{0, 0, 0},
			},
			wantStatus: domain.StatusInProgress,
			wantOver:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, gotOver := CheckGameOver(tt.board)
			if gotStatus != tt.wantStatus || gotOver != tt.wantOver {
				t.Errorf("CheckGameOver() = (%v, %v), want (%v, %v)", gotStatus, gotOver, tt.wantStatus, tt.wantOver)
			}
		})
	}
}

func TestValidateBoard(t *testing.T) {
	oldGame := domain.Game{
		ID: uuid.New(),
		Board: domain.Board{
			{1, 0, 0},
			{0, 2, 0},
			{0, 0, 0},
		},
		Status: domain.StatusInProgress,
	}

	// Valid move
	validBoard := domain.Board{
		{1, 1, 0},
		{0, 2, 0},
		{0, 0, 0},
	}
	if err := ValidateBoard(oldGame, validBoard); err != nil {
		t.Errorf("ValidateBoard() valid move failed: %v", err)
	}

	// Modified previous move
	modifiedBoard := domain.Board{
		{0, 0, 0},
		{0, 2, 0},
		{1, 0, 0},
	}
	if err := ValidateBoard(oldGame, modifiedBoard); err != domain.ErrPreviousMovesModified {
		t.Errorf("expected ErrPreviousMovesModified, got %v", err)
	}

	// Two moves in one turn
	twoMovesBoard := domain.Board{
		{1, 1, 1},
		{0, 2, 0},
		{0, 0, 0},
	}
	if err := ValidateBoard(oldGame, twoMovesBoard); err != domain.ErrInvalidMoveCount {
		t.Errorf("expected ErrInvalidMoveCount, got %v", err)
	}
}

func TestMinimaxBlockWin(t *testing.T) {
	// X is about to win on row 0: {1, 1, 0}
	// O must block at (0, 2)
	board := domain.Board{
		{1, 1, 0},
		{0, 2, 0},
		{0, 0, 0},
	}

	move, err := FindBestMove(board)
	if err != nil {
		t.Fatalf("FindBestMove() unexpected error: %v", err)
	}

	if move.Row != 0 || move.Col != 2 {
		t.Errorf("FindBestMove() expected blocking move (0, 2), got (%d, %d)", move.Row, move.Col)
	}
}

func TestMinimaxTakeWinningMove(t *testing.T) {
	// O can win at (1, 2):
	// {1, 1, 0}
	// {2, 2, 0} -> (1, 2) wins!
	board := domain.Board{
		{1, 1, 0},
		{2, 2, 0},
		{1, 0, 0},
	}

	move, err := FindBestMove(board)
	if err != nil {
		t.Fatalf("FindBestMove() unexpected error: %v", err)
	}

	if move.Row != 1 || move.Col != 2 {
		t.Errorf("FindBestMove() expected winning move (1, 2), got (%d, %d)", move.Row, move.Col)
	}
}
