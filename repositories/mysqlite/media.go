package mysqlite

import (
	"context"
	"fmt"
	"github.com/Slava02/helperBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *mysqlite) Save(ctx context.Context, update tgbotapi.Update, user *models.User, mediaType models.Media) error {
	q := `INSERT INTO media (telegram_id, media, media_type) VALUES (?, ?, ?)`
	_, err := m.db.ExecContext(ctx, q, user.TelegramID, update.Message.Text, mediaType)
	if err != nil {
		m.logger.Error(err)
		return err
	}

	return nil
}

func (m *mysqlite) PickRandom(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, mediaType models.Media) (string, error) {
	q := `SELECT media FROM media WHERE telegram_id = ? AND media_type = ? ORDER BY RANDOM() LIMIT 1`

	var media string
	row := m.db.QueryRow(q, user.TelegramID, mediaType)
	if err := row.Scan(&media); err != nil {
		return "", fmt.Errorf("can't pickRandom: %w", err)
	}

	return media, nil
}
