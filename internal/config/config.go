package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ppaprikaa/shorty/internal/env"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"config",
	fx.Provide(MustInit),
	fx.Provide(GetENV),
	fx.Provide(HttpServerFromConfig),
	fx.Provide(MainStorageFromConfig),
	fx.Provide(RefreshTokensStorageFromConfig),
	fx.Provide(MailerFromConfig),
)

type Config struct {
	Env                  env.ENV               `toml:"env" env:"SHORTY_ENV"`
	HttpServer           *HttpServer           `toml:"http_server"`
	MainStorage          *MainStorage          `toml:"main_storage"`
	RefreshTokensStorage *RefreshTokensStorage `toml:"refresh_tokens_storage"`
	Mailer               *Mailer               `toml:"mailer"`
}

func GetENV(cfg *Config) env.ENV { return cfg.Env }

func Init() (*Config, error) {
	var (
		configFilePath string
		cfg            Config
	)

	if configFilePath = os.Getenv("SHORTY_CONFIG_PATH"); configFilePath == "" {
		configFilePath = "config/config.toml"
	}

	if err := cleanenv.ReadConfig(configFilePath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func MustInit() *Config {
	cfg, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
