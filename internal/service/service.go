package service

import (
	"context"

	"github.com/google/uuid"

	"tictactoe/internal/domain"
)

type gameService struct {
	repo domain.GameRepository
}

// NewGameService создает новый экземпляр GameService
func NewGameService(repo domain.GameRepository) domain.GameService {
	return &gameService{
		repo: repo,
	}
}

// GetNextMove вычисляет следующий оптимальный ход для текущей игры алгоритмом Минимакс
func (s *gameService) GetNextMove(game domain.Game) (domain.Move, error) {
	return FindBestMove(game.Board)
}

// ValidateBoard проверяет, что не изменены предыдущие ходы и сделан ровно один ход пользователем
func (s *gameService) ValidateBoard(oldGame domain.Game, newBoard domain.Board) error {
	return ValidateBoard(oldGame, newBoard)
}

// CheckGameOver проверяет окончание игры
func (s *gameService) CheckGameOver(board domain.Board) (domain.GameStatus, bool) {
	return CheckGameOver(board)
}

// ProcessTurn обрабатывает ход игрока, делает ответный ход AI и сохраняет игру
func (s *gameService) ProcessTurn(ctx context.Context, gameID uuid.UUID, newBoard domain.Board) (domain.Game, error) {
	oldGame, err := s.repo.Get(ctx, gameID)
	if err != nil {
		return domain.Game{}, err
	}

	// 1. Валидация хода пользователя
	if err := s.ValidateBoard(oldGame, newBoard); err != nil {
		return domain.Game{}, err
	}

	currentGame := domain.Game{
		ID:     oldGame.ID,
		Board:  newBoard,
		Status: domain.StatusInProgress,
	}

	// 2. Проверка окончания игры после хода пользователя (победа игрока или ничья)
	status, isOver := s.CheckGameOver(currentGame.Board)
	if isOver {
		currentGame.Status = status
		if err := s.repo.Save(ctx, currentGame); err != nil {
			return domain.Game{}, err
		}
		return currentGame, nil
	}

	// 3. Расчет хода компьютера алгоритмом Минимакс
	aiMove, err := s.GetNextMove(currentGame)
	if err != nil {
		return domain.Game{}, err
	}

	// 4. Применение хода компьютера
	currentGame.Board[aiMove.Row][aiMove.Col] = domain.CellO

	// 5. Проверка окончания игры после хода компьютера
	statusAfterAI, _ := s.CheckGameOver(currentGame.Board)
	currentGame.Status = statusAfterAI

	// 6. Сохранение обновленного состояния в репозиторий
	if err := s.repo.Save(ctx, currentGame); err != nil {
		return domain.Game{}, err
	}

	return currentGame, nil
}

// CreateGame создает новую игру с пустым полем и сохраняет ее
func (s *gameService) CreateGame(ctx context.Context) (domain.Game, error) {
	newGame := domain.Game{
		ID:     uuid.New(),
		Board:  domain.Board{},
		Status: domain.StatusInProgress,
	}

	if err := s.repo.Save(ctx, newGame); err != nil {
		return domain.Game{}, err
	}

	return newGame, nil
}

// GetGame возвращает игру по её идентификатору
func (s *gameService) GetGame(ctx context.Context, gameID uuid.UUID) (domain.Game, error) {
	return s.repo.Get(ctx, gameID)
}
