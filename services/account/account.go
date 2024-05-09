// package account implements account functionalities.
package account

import (
	"github.com/Slava02/helperBot/pkg/logger"
	"github.com/Slava02/helperBot/repositories"
	"github.com/Slava02/helperBot/services"
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
