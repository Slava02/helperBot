// Package bot handles bot activities.
package bot

import (
	"context"
	"github.com/Slava02/helperBot/models"
	"regexp"

	_ "github.com/Slava02/helperBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mediaType models.Media
var removePat = regexp.MustCompile(`^r#.`)

// Callback handles callback requests.
func (h *handler) Callback(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	deleteLastMessage(tgbot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)

	user := &models.User{
		TelegramID: update.CallbackQuery.From.ID,
		Username:   update.CallbackQuery.From.UserName,
		FirstName:  update.CallbackQuery.From.FirstName,
		LastName:   update.CallbackQuery.From.LastName,
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	var err error
	if err != nil {
		h.logger.Error(err)
	}

	switch data := update.CallbackQuery.Data; {
	case data == "information":
		msg = h.callback.Information(ctx, msg, user)
	case data == "help":
		msg = h.callback.Help(ctx, msg, user)
	case data == "media":
		msg = h.callback.Media(ctx, msg, user)
	case data == "book":
		mediaType = models.Book
		msg = h.callback.MediaOptions(ctx, msg, user)
	case data == "movie":
		mediaType = models.Movie
		msg = h.callback.MediaOptions(ctx, msg, user)
	case data == "saveWhat":
		UserEnterMode = true
		msg = h.callback.SaveWhat(ctx, msg, user, mediaType)
	case data == "pickRandom":
		msg = h.callback.PickRandom(ctx, msg, user, mediaType)
	case data == "save":
		msg = h.callback.Save(ctx, msg, user, mediaType)
	case data == "removeWhat":
		msg = h.callback.RemoveWhat(ctx, msg, user, mediaType)
	case removePat.MatchString(data):
		msg = h.callback.Remove(ctx, msg, user, mediaType, data)
	case data == "list":
		msg = h.callback.List(ctx, msg, user, mediaType)
	case data == "backToMain":
		UserEnterMode = false
		msg = h.callback.Menu(ctx, msg, user)
	}

	_, err = tgbot.Send(msg)
	if err != nil {
		h.logger.Error(err)
	}
}
