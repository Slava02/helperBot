package mysqlite

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Slava02/helperBot/configs"
	"github.com/Slava02/helperBot/pkg/logger"
	"github.com/Slava02/helperBot/repositories"
	_ "github.com/mattn/go-sqlite3"
)

type mysqlite struct {
	db     *sql.DB
	logger logger.Logger
}

func New(ctx context.Context, cfg configs.MySQLite, logger logger.Logger) (repositories.DB, error) {
	db, err := sql.Open("sqlite3", path(cfg))
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("can not ping the database: %v", err)
	}

	return &mysqlite{db: db, logger: logger}, nil
}

func (m *mysqlite) Close(ctx context.Context) error {
	err := m.db.Close()
	if err != nil {
		return fmt.Errorf("can not close the database: %v", err)
	}

	return nil
}

func path(cfg configs.MySQLite) string {
	return fmt.Sprintf("%s/%s", cfg.Path, cfg.DBName)
}
