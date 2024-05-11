package callback

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/keyboards"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
	"strings"
)

func (c *call) SaveWhat(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = messages.SaveWhat()
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}

func (c *call) Save(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = messages.SaveSuccess()
	return msg
}

func (c *call) RemoveWhat(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	mediaObjects, err := c.db.List(ctx, msg, user, media)
	if err != nil {
		c.logger.Error(err)
		msg.Text = messages.Err()
	} else {
		msg.Text = messages.RemoveWhat()
	}
	msg.ReplyMarkup = keyboards.RemoveWhat(mediaObjects)

	return msg
}

func (c *call) Remove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media, callBackData string) tgbotapi.MessageConfig {
	err := c.db.Remove(ctx, msg, user, media, callBackDataToName(callBackData))

	if err != nil {
		c.logger.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			msg.Text = messages.NoData()
		} else {
			msg.Text = messages.Err()
		}
	} else {
		msg.Text = messages.RemoveSuccess()
	}

	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}

func callBackDataToName(data string) string {
	re, _ := regexp.Compile(`\#.*`)
	return strings.Replace(re.FindString(data), "#", "", 1)
}
