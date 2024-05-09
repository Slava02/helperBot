// Package services implements all aplication services.
package services

import (
	"context"

	"github.com/Slava02/helperBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Callback implements callback functionalities.
type Callback interface {
	Menu(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	Help(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	Information(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	Media(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	SaveOrRemove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	SaveWhat(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig
	Save(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig
	Remove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig
}
