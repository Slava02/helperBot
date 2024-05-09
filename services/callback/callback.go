// package callback implements callback request functionalities.
package callback

import (
	"github.com/Slava02/helperBot/pkg/logger"
	"github.com/Slava02/helperBot/repositories"
	"github.com/Slava02/helperBot/services"
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
