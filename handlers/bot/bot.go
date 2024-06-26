// Package bot handles bot activities.
package bot

import (
	"context"
	"fmt"
	"github.com/Slava02/helperBot/configs"
	"github.com/Slava02/helperBot/handlers"
	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/logger"
	"github.com/Slava02/helperBot/repositories"
	"github.com/Slava02/helperBot/services/account"
	"github.com/Slava02/helperBot/services/callback"
	"github.com/Slava02/helperBot/services/command"
	"github.com/Slava02/helperBot/services/message"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var UserEnterMode models.EnterMode

type bot struct {
	handler *handler
	logger  logger.Logger
}

// New creates a new bot.
func New(db repositories.DB, logger logger.Logger) handlers.Bot {
	acc := account.New(db, logger)
	msg := message.New(db, logger)
	cmd := command.New(db, logger)
	call := callback.New(db, logger)

	handler := handler{
		logger:   logger,
		account:  acc,
		message:  msg,
		command:  cmd,
		callback: call,
	}

	return &bot{handler: &handler, logger: logger}
}

// Start starts the bot.
func (b *bot) Start(ctx context.Context, cfg configs.App, debugMode bool) error {
	b.logger.Info("Bot Starting...")

	tgbot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		b.logger.Error(fmt.Errorf("can't connect to the bot: %v", err))

		return err
	}

	tgbot.Debug = debugMode

	b.logger.Info(fmt.Sprintf("Authorized on account %s", tgbot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgbot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			go b.handler.Command(ctx, tgbot, update)
		} else if update.CallbackQuery != nil {
			go b.handler.Callback(ctx, tgbot, update)
		} else if update.Message != nil {
			go b.handler.Message(ctx, tgbot, update)
		} else {
			continue
		}
	}

	return nil
}
