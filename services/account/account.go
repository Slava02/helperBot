// package account implements account functionalities.
package account

import (
	"cleanBot/pkg/logger"
	"cleanBot/repositories"
	"cleanBot/services"
)

type acc struct {
	db     repositories.DB
	logger logger.Logger
}

// New returns a new user account service.
func New(db repositories.DB, logger logger.Logger) services.Account {
	return &acc{
		db:     db,
		logger: logger,
	}
}
