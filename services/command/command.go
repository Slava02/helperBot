// package command implements command request functionalities.
package command

import (
	"cleanBot/pkg/logger"
	"cleanBot/repositories"
	"cleanBot/services"
)

type cmd struct {
	db     repositories.DB
	logger logger.Logger
}

// New returns a command service.
func New(db repositories.DB, logger logger.Logger) services.Command {
	return &cmd{
		db:     db,
		logger: logger,
	}
}
