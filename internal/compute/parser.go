package compute

import (
	"strings"

	"go.uber.org/zap"
)

type Parser struct {
	logger *zap.Logger
}

func NewParser(logger *zap.Logger) *Parser {
	return &Parser{
		logger: logger,
	}
}

func (p *Parser) Parse(queryStr string) (*Query, error) {
	tokens := strings.Fields(queryStr)
	if len(tokens) == 0 {
		p.logger.Debug("empty tokens", zap.String("query", queryStr))

		return &Query{}, errEmptyQuery
	}

	command := tokens[0]
	commandID := mapCommandNameToCommandID(command)
	if commandID == UnknownCommandID {
		p.logger.Debug("unknown command", zap.String("query", command))

		return &Query{}, errInvalidCommand
	}

	query := newQuery(commandID, tokens[1:])
	argumentsNumber := getArgumentsNumberByCommandID(commandID)
	if len(query.Arguments()) != argumentsNumber {
		p.logger.Debug("invalid arguments for query", zap.String("query", queryStr))

		return &Query{}, errInvalidArgs
	}

	return query, nil
}
