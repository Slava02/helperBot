// Package cmd contains main functionality for the application using CLI commands.
package cmd

import (
	"cleanBot/repositories/mysqlite"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cleanBot/configs"
	"cleanBot/handlers/bot"
	"cleanBot/models"
	"cleanBot/pkg/logger/zap"
	"cleanBot/repositories"
	"cleanBot/repositories/mysql"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
)

var (
	serveCMD = &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Serving telegram bot",
		Action:  serve,
		Subcommands: []*cli.Command{
			debugCMD,
		},
	}

	debugCMD = &cli.Command{
		Name:    "debug",
		Aliases: []string{"d"},
		Usage:   "Turn on debug mode",
		After:   serve,
		Action:  debug,
	}

	debugMode = false
)

// serve is the main command that runs the application.
func serve(c *cli.Context) error {
	err := os.Mkdir("logs", os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	file, err := os.OpenFile("logs/bot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	logger := zap.New(file, zapcore.InfoLevel)

	cfg, err := configs.New(logger, models.Development)
	if err != nil {
		return err
	}

	ctx := context.TODO()

	var db repositories.DB

	switch cfg.App.Driver {
	case "mysql":
		db, err = mysql.New(ctx, cfg.MySQL, logger)
		if err != nil {
			logger.Error(err)

			return err
		}
	case "mysqlite":
		db, err = mysqlite.New(ctx, cfg.MySQLite, logger)
		if err != nil {
			logger.Error(err)

			return err
		}
	}

	defer db.Close(ctx)

	bot := bot.New(db, logger)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		err = bot.Start(ctx, cfg.App, debugMode)
		if err != nil {
			logger.Error(fmt.Sprintf("bot failed to start: %v", err))
			os.Exit(1)
		}
	}()

	<-sigChan

	logger.Info("Received an interrupt, Bot stopped...")

	return nil
}

// debug is a subcommand of the serve command that turns debug mode on.
func debug(c *cli.Context) error {
	log.Println("Debug mode is on")

	debugMode = true

	return nil
}
