package callback

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/keyboards"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *call) MediaOptions(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.MediaOptions()
	msg.ReplyMarkup = keyboards.MediaOptions()

	return msg
}

func (c *call) SaveWhat(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = messages.SaveWhat()
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}

func (c *call) SaveSuccess(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = messages.SaveSuccess()

	return msg
}

func (c *call) Remove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = messages.Remove()
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}

func (c *call) PickRandom(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, mediaType models.Media) tgbotapi.MessageConfig {
	var err error
	msg.Text, err = c.db.PickRandom(ctx, msg, user, mediaType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			msg.Text = messages.NoData()
		} else {
			msg.Text = messages.PickRandomErr()
		}
	}
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}
