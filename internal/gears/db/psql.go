package db

import (
	"context"
	"shorty/internal/gears/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func ConnectPSQL(ctx context.Context, cfg *config.PsqlConfig) (*pgx.Conn, error) {
	connCfg := &pgx.ConnConfig{
		Config: pgconn.Config{
			Host:     cfg.Host,
			Port:     cfg.Port,
			Database: cfg.DB,
			User:     cfg.User,
			Password: cfg.Password,
		},
	}

	return pgx.ConnectConfig(ctx, connCfg)
}
