package storage

import "go.uber.org/zap"

type StorageEngine interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Storage struct {
	engine StorageEngine
	logger *zap.Logger
}

func NewStorage(engine StorageEngine, logger *zap.Logger) *Storage {
	return &Storage{
		engine: engine,
		logger: logger,
	}
}

func (s *Storage) Set(key string, value string) error {
	err := s.engine.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Get(key string) (string, error) {
	value, err := s.engine.Get(key)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (s *Storage) Delete(key string) error {
	err := s.engine.Delete(key)
	if err != nil {
		return err
	}

	return nil
}
