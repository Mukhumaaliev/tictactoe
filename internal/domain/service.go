package domain

import (
	"context"
	"errors"
)

type Service struct {
	repo GameRepository
}

func NewService(repo GameRepository) GameService {
	return &Service{
		repo: repo,
	}
}

type GameService interface {
	CreateGame(ctx context.Context) (*Game, error)
	NextMove(ctx context.Context, uuid string, newGame *Game) (*Game, error)
}

type GameRepository interface {
	Save(ctx context.Context, game *Game) error
	Get(ctx context.Context, uuid string) (*Game, error)
}

func (s *Service) CreateGame(ctx context.Context) (*Game, error) {
	game := NewGame()
	if err := s.repo.Save(ctx, game); err != nil {
		return nil, err
	}
	return game, nil
}

func (s *Service) NextMove(ctx context.Context, uuid string, newGame *Game) (*Game, error) {
	game, err := s.repo.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if err := isValidBoard(game.Board, newGame.Board); err != nil {
		return nil, err
	}
	game.Board = newGame.Board
	updateGameStatus(game)

	if game.Status != GameStatusInProgress {
		return game, errors.New("Игра закончилась!")
	}
	makeRobotMove(game)
	s.repo.Save(ctx, game)
	if game.Status != GameStatusInProgress {
		return game, errors.New("Игра закончилась!")
	}
	return game, nil
}
