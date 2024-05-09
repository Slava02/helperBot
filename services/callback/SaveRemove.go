package callback

import (
	"context"
	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/keyboards"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *call) SaveOrRemove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.SaveOrRemove()
	msg.ReplyMarkup = keyboards.SaveOrRemove()

	return msg
}

func (c *call) SaveWhat(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = messages.SaveWhat()
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}

// TODO разобраться что делают эти функции
func (c *call) Save(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = messages.SaveSuccess()

	return msg
}

func (c *call) Remove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = messages.Remove()
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}
