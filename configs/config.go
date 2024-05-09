// Package configs implement configurations that the application needs.
package configs

import (
	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/logger"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		MySQL    MySQL
		MySQLite MySQLite
	}

	App struct {
		Name   string `env:"BOT_NAME,required"`
		Token  string `env:"BOT_APITOKEN,required"`
		Driver string `env:"BOT_DB_DRIVER,required"`
	}

	MySQL struct {
		Username string `env:"BOT_MYSQL_USERNAME,required"`
		Password string `env:"BOT_MYSQL_PASSWORD,required"`
		Protocol string `env:"BOT_MYSQL_PROTOCOL,required"`
		Host     string `env:"BOT_MYSQL_HOST,required"`
		Port     string `env:"BOT_MYSQL_PORT,required"`
		DBName   string `env:"BOT_MYSQL_DBNAME,required"`
	}

	MySQLite struct {
		Path   string `env:"BOT_MYSQLite_PATH,required"`
		DBName string `env:"BOT_MYSQLite_DBNAME,required"`
	}
)

// New returns the config, if can't open .env file or parse environment variables it returns an error.
// TODO: разобраться как работают прод и дев режимы
// make sure to delete the .env file and pass production mode in production.
func New(logger logger.Logger, mode models.Mode) (*Config, error) {
	if mode == models.Development {
		err := godotenv.Load("configs/.env")
		if err != nil {
			logger.Error(err)

			return nil, err
		}
	}

	cfg := Config{}

	err := env.Parse(&cfg)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return &cfg, nil
}
