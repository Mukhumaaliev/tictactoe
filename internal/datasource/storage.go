package datasource

import (
	"errors"
	"sync"
)

type Storage struct {
	games sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) saveGame(g *Game) {
	s.games.Store(g.UUID, g)
}

func (s *Storage) getGame(uuid string) (*Game, error) {
	val, ok := s.games.Load(uuid)
	if !ok {
		return &Game{}, errors.New("Игры с таким айди не существует!")
	}
	game := val.(*Game)
	return game, nil
}
