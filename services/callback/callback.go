// package callback implements callback request functionalities.
package callback

import (
	"cleanBot/pkg/logger"
	"cleanBot/repositories"
	"cleanBot/services"
)

type call struct {
	db     repositories.DB
	logger logger.Logger
}

// New returns a callback service.
func New(db repositories.DB, logger logger.Logger) services.Callback {
	return &call{
		db:     db,
		logger: logger,
	}
}
