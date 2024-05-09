// Package repositories implements database functionalities.
package repositories

import (
	"context"

	"github.com/Slava02/helperBot/models"
)

// DB implements all database functionalities.
type DB interface {
	User
	Close(ctx context.Context) error
}

// User implements user functionalities.
type User interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, telegramID int64) (*models.User, error)
}
