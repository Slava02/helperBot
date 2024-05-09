// Package repositories implements database functionalities.
package repositories

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Slava02/helperBot/models"
)

// DB implements all database functionalities.
type DB interface {
	User
	Media
	Close(ctx context.Context) error
}

// User implements user functionalities.
type User interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, telegramID int64) (*models.User, error)
}

type Media interface {
	Save(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, mediaType models.Media) error
}
