// Package services implements all aplication services.
package services

import (
	"cleanBot/models"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Command implements command functionalities.
type Command interface {
	Help(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
}
