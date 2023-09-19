package mongo

import (
	"context"

	"github.com/ppaprikaa/shorty/internal/config"
	"github.com/ppaprikaa/shorty/internal/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"db.mongo",
	fx.Provide(Connect),
	fx.Invoke(func(*DB) {}),
)

type DB struct {
	Client *mongo.Client
}

func Connect(ctx context.Context, lc fx.Lifecycle, cfg *config.MainStorage) (*DB, error) {
	var (
		DB  = &DB{}
		err error
	)

	lc.Append(fx.Hook{
		OnStart: func(startCtx context.Context) error {
			DB.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.DSN))
			if err != nil {
				return err
			}
			log.FromContext(ctx).Info("MONGO LOADED")
			return nil
		},
		OnStop: func(stopCtx context.Context) error {
			log.FromContext(ctx).Info("MONGO CLOSED")
			return DB.Close(stopCtx)
		},
	})

	return DB, nil
}

func (d *DB) Close(ctx context.Context) error {
	return d.Client.Disconnect(ctx)
}
