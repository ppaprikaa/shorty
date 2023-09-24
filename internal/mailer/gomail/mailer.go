package gomail

import (
	"context"

	"github.com/ppaprikaa/shorty/internal/config"
	"github.com/ppaprikaa/shorty/internal/log"
	"go.uber.org/fx"
	"gopkg.in/gomail.v2"
)

var Module = fx.Module(
	"mailer.gomail",
	fx.Provide(Connect),
	fx.Invoke(func(*Mailer) {}),
)

type Mailer struct {
	Client gomail.SendCloser
}

func Connect(ctx context.Context, lc fx.Lifecycle, cfg *config.Mailer) (*Mailer, error) {
	var (
		Mailer = new(Mailer)
		err    error
	)

	lc.Append(fx.Hook{
		OnStart: func(startCtx context.Context) error {
			Mailer.Client, err = gomail.NewDialer(
				cfg.Host,
				cfg.Port,
				cfg.Username,
				cfg.Password,
			).Dial()
			if err != nil {
				return err
			}
			log.FromContext(ctx).Info("MAILER LOADED")

			return nil
		},
		OnStop: func(stopCtx context.Context) error {
			log.FromContext(ctx).Info("MAILER CLOSED")
			return Mailer.Close(stopCtx)
		},
	})

	return Mailer, nil
}

func (m *Mailer) Close(ctx context.Context) error {
	return m.Client.Close()
}
