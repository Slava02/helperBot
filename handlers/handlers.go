// package handlers handles all incoming requests.
package handlers

import (
	"context"

	"github.com/Slava02/helperBot/configs"
)

// Bot implements bot functionalities.
type Bot interface {
	Start(ctx context.Context, cfg configs.App, debugMode bool) error
}
