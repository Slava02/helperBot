// Package services implements all aplication services.
package services

import (
	"context"

	"github.com/Slava02/helperBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message implements message functionalities.
type Message interface {
	Wrong(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	Save(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, mediaType models.Media) (tgbotapi.MessageConfig, error)
}
