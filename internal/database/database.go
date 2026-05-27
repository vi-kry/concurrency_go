package database

import (
	"errors"

	"go.uber.org/zap"

	"github.com/vi-kry/concurrency_go/internal/compute"
)

const (
	operationStatusOk   = "[ok]"
	operationStatusFail = "[fail]"
)

type Compute interface {
	Parse(queryStr string) (*compute.Query, error)
}

type Storage interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Database struct {
	compute Compute
	storage Storage
	logger  *zap.Logger
}

func NewDatabase(
	compute Compute,
	storage Storage,
	logger *zap.Logger,
) *Database {
	return &Database{
		compute: compute,
		storage: storage,
		logger:  logger,
	}
}

func (d *Database) HandleQuery(
	queryStr string,
) (string, error) {
	d.logger.Debug("handling query", zap.String("query", queryStr))
	query, err := d.compute.Parse(queryStr)
	if err != nil {
		return "", err
	}

	switch query.CommandID() {
	case compute.SetCommandID:
		handleErr := d.handleSetQuery(query)
		if handleErr != nil {
			return "", handleErr
		}

		return operationStatusOk, nil
	case compute.GetCommandID:
		value, handleErr := d.handleGetQuery(query)
		if handleErr != nil {
			return "", handleErr
		}

		return value, nil
	case compute.DeleteCommandID:
		handleErr := d.handleDeleteQuery(query)
		if handleErr != nil {
			return "", handleErr
		}

		return operationStatusOk, nil
	}

	d.logger.Error("compute layer is incorrect", zap.Int("command_id", query.CommandID()))

	return "", errors.New(operationStatusFail)
}

func (d *Database) handleSetQuery(query *compute.Query) error {
	args := query.Arguments()

	err := d.storage.Set(args[0], args[1])
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) handleGetQuery(query *compute.Query) (string, error) {
	args := query.Arguments()

	value, err := d.storage.Get(args[0])
	if err != nil {
		return "", err
	}

	return value, nil
}

func (d *Database) handleDeleteQuery(query *compute.Query) error {
	args := query.Arguments()

	err := d.storage.Delete(args[0])
	if err != nil {
		return err
	}

	return nil
}
