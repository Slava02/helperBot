// Package mysql implements mysql database functionalities.
package mysql

import (
	"context"
	"fmt"

	"github.com/Slava02/helperBot/models"
)

// CreateUser creates a new user.
func (m *mysql) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (telegram_id, username, firstname, lastname, joined_at) VALUES(?, ?, ?, ?, ?)"

	_, err := m.db.ExecContext(ctx, query, user.TelegramID, user.Username, user.FirstName, user.LastName, user.JoinedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetUser returns a user.
func (m *mysql) GetUser(ctx context.Context, telegramID int64) (*models.User, error) {
	user := models.User{}

	query := "SELECT * FROM users WHERE telegram_id=?"

	row := m.db.QueryRowContext(ctx, query, telegramID)
	err := row.Scan(&user.ID, &user.TelegramID, &user.Username, &user.FirstName, &user.LastName, &user.JoinedAt)
	if err != nil {
		return nil, fmt.Errorf("can not fetch user %d from database: %v", telegramID, err)
	}

	return &user, err
}
