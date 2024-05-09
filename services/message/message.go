// package account implements message request functionalities.
package message

import (
	"github.com/Slava02/helperBot/pkg/logger"
	"github.com/Slava02/helperBot/repositories"
	"github.com/Slava02/helperBot/services"
)

type msg struct {
	db     repositories.DB
	logger logger.Logger
}

// New returns a message service.
func New(db repositories.DB, logger logger.Logger) services.Message {
	return &msg{
		db:     db,
		logger: logger,
	}
}
