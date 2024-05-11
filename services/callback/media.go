package callback

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/keyboards"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *call) Media(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Media(user)
	msg.ReplyMarkup = keyboards.Media()

	return msg
}

func (c *call) MediaOptions(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.MediaOptions()
	msg.ReplyMarkup = keyboards.MediaOptions()

	return msg
}

func (c *call) PickRandom(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, mediaType models.Media) tgbotapi.MessageConfig {
	var err error
	msg.Text, err = c.db.PickRandom(ctx, msg, user, mediaType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			msg.Text = messages.NoData()
		} else {
			msg.Text = messages.Err()
		}
	}
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}

func (c *call) List(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	mediaObjects, err := c.db.List(ctx, msg, user, media)
	if err != nil {
		c.logger.Error(err)
		msg.Text = messages.Err()
	} else if len(mediaObjects) == 0 {
		msg.Text = messages.NoData()
	} else {
		msg.Text = mediaToMsg(mediaObjects)
	}

	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}

func mediaToMsg(media []string) string {
	var res string
	for i, m := range media {
		if i > 0 {
			res += strconv.Itoa(i) + ". " + m + "\n"
		}
	}
	return res
}
