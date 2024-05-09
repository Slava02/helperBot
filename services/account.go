// Package services implements all aplication services.
package services

import (
	"context"

	"github.com/Slava02/helperBot/models"
)

// Account implements user account functionalities.
type Account interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, user *models.User) (*models.User, error)
	IsExist(ctx context.Context, user *models.User) bool
}
