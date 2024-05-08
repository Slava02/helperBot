// Package services implements all aplication services.
package services

import (
	"context"

	"cleanBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message implements message functionalities.
type Message interface {
	Wrong(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
}
