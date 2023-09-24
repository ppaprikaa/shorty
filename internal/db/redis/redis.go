package redis

import (
	"context"

	"github.com/ppaprikaa/shorty/internal/config"
	"github.com/ppaprikaa/shorty/internal/log"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"db.redis",
	fx.Provide(Connect),
	fx.Invoke(func(*DB) {}),
)

type DB struct {
	Client *redis.Client
}

func Connect(ctx context.Context, lc fx.Lifecycle, cfg *config.RefreshTokensStorage) (*DB, error) {
	var (
		DB  = new(DB)
		err error
	)

	lc.Append(fx.Hook{
		OnStart: func(startCtx context.Context) error {
			client := redis.NewClient(
				&redis.Options{
					Addr:     cfg.DSN,
					Password: cfg.Password,
					DB:       cfg.DB,
				},
			)

			if err = client.Ping(startCtx).Err(); err != nil {
				return err
			}
			DB.Client = client
			log.FromContext(ctx).Info("REDIS LOADED")

			return nil
		},
		OnStop: func(stopCtx context.Context) error {
			log.FromContext(ctx).Info("REDIS CLOSED")
			return DB.Close()
		},
	})

	return DB, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}
