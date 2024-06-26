// Package bot handles bot activities.
package bot

import (
	"github.com/Slava02/helperBot/pkg/logger"
	"github.com/Slava02/helperBot/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handler handles all services.
type handler struct {
	logger   logger.Logger
	account  services.Account
	message  services.Message
	command  services.Command
	callback services.Callback
}

// deleteLastMessage deletes recieved request for user.
func deleteLastMessage(tgbot *tgbotapi.BotAPI, chatID int64, messageID int) {
	deleteLastMessage := tgbotapi.NewDeleteMessage(chatID, messageID)

	tgbot.Request(deleteLastMessage)
}
