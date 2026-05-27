package storage

import (
	"go.uber.org/zap"
)

type Engine struct {
	data   map[string]string
	logger *zap.Logger
}

func NewEngine(logger *zap.Logger) *Engine {
	return &Engine{
		data:   make(map[string]string),
		logger: logger,
	}
}

func (e *Engine) Set(key, value string) error {
	e.data[key] = value

	return nil
}

func (e *Engine) Get(key string) (string, error) {
	if value, ok := e.data[key]; ok {
		return value, nil
	}

	return "", errKeyNotFound
}

func (e *Engine) Delete(key string) error {
	if _, ok := e.data[key]; ok {
		delete(e.data, key)

		return nil
	}

	return errKeyNotFound
}
