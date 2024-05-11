package message

import (
	"context"
	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (m *msg) Save(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, mediaType models.Media, update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	isExist, err := m.db.IsExist(ctx, update, user, mediaType)
	if err != nil {
		msg.Text = messages.SaveError()
	} else if isExist {
		msg.Text = messages.AlreadySaved() // TODO
	} else {
		err = m.db.Save(ctx, update, user, mediaType)
		if err != nil {
			m.logger.Error(err)
			msg.Text = messages.SaveError()
			return msg, err
		} else {
			msg.Text = messages.SaveSuccess()
		}
	}

	return msg, nil
}
