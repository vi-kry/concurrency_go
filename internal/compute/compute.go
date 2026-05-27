package compute

import "go.uber.org/zap"

type Compute struct {
	parser *Parser
	logger *zap.Logger
}

func NewCompute(parser *Parser, logger *zap.Logger) *Compute {
	return &Compute{
		parser: parser,
		logger: logger,
	}
}

func (c *Compute) Parse(queryStr string) (*Query, error) {
	return c.parser.Parse(queryStr)
}
