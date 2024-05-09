// package command implements command request functionalities.
package command

import (
	"github.com/Slava02/helperBot/pkg/logger"
	"github.com/Slava02/helperBot/repositories"
	"github.com/Slava02/helperBot/services"
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
