package mysqlite

import (
	"context"
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
