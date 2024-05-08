// package account implements message request functionalities.
package message

import (
	"cleanBot/pkg/logger"
	"cleanBot/repositories"
	"cleanBot/services"
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
