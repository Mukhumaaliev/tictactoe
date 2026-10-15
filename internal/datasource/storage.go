package datasource

import "sync"

// Storage представляет собой потокобезопасное in-memory хранилище на базе sync.Map
type Storage struct {
	games sync.Map
}

// NewStorage создает новый экземпляр Storage
func NewStorage() *Storage {
	return &Storage{}
}

// Set сохраняет сущность игры по идентификатору
func (s *Storage) Set(id string, entity GameEntity) {
	s.games.Store(id, entity)
}

// Get возвращает сущность игры по идентификатору, если она существует
func (s *Storage) Get(id string) (GameEntity, bool) {
	val, ok := s.games.Load(id)
	if !ok {
		return GameEntity{}, false
	}
	entity, ok := val.(GameEntity)
	return entity, ok
}

// Delete удаляет игру по идентификатору
func (s *Storage) Delete(id string) {
	s.games.Delete(id)
}
