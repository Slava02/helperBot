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

func (m *mysqlite) IsExist(ctx context.Context, update tgbotapi.Update, user *models.User, mediaType models.Media) (bool, error) {
	q := `SELECT COUNT(*) FROM media WHERE media_type = ? AND media = ?`

	var count int

	if err := m.db.QueryRow(q, mediaType, update.Message.Text).Scan(&count); err != nil {
		m.logger.Error(err)
		return false, fmt.Errorf("couldn't check if Exists: %w", err)
	}

	return count > 0, nil
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

func (m *mysqlite) List(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, mediaType models.Media) ([]string, error) {
	res := make([]string, 0)

	q := `SELECT media FROM media WHERE media_type = ?`

	rows, err := m.db.Query(q, mediaType)
	if err != nil {
		m.logger.Error(err)
		return nil, fmt.Errorf("couldn't excecute List: %w", err)
	}
	defer rows.Columns()

	for rows.Next() {
		var tmp string
		if err = rows.Scan(&tmp); err != nil {
			return nil, fmt.Errorf("couldn't scan List: %w", err)
		}

		res = append(res, tmp)
	}

	return res, nil
}

func (m *mysqlite) Remove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, mediaType models.Media, name string) error {
	q := `DELETE FROM media WHERE media_type = ? AND media = ?`

	if _, err := m.db.Exec(q, mediaType, name); err != nil {
		return fmt.Errorf("couldn't remove object: %w", err)
	}

	return nil
}
