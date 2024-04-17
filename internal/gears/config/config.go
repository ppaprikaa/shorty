package config

import (
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "LOCAL"
	EnvDev   = "DEV"
	EnvProd  = "Prod"
)

func Load() *Config {
	var cfg *Config = new(Config)

	path := strings.TrimSpace(os.Getenv("SHORTY_CONFIG"))
	if path == "" {
		path = "config/config.json"
	}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic(err)
	}

	return cfg
}

type Config struct {
	HttpServer HttpConfig `json:"http-server"`
	Psql       PsqlConfig `json:"psql"`
	Mailer     Mailer     `json:"mailer"`
	Tokens     Tokens     `json:"tokens"`
}

type HttpConfig struct {
	Host string `json:"host" env:"SHORTY_HOST" env-default:""`
	Port uint16 `json:"port" env:"SHORTY_PORT" env-default:"56565"`
}

type Mailer struct {
	Host     string `json:"host" env:"SHORTY_SMTP_HOST"`
	Port     int    `json:"port" env:"SHORTY_SMTP_PORT"`
	Username string `json:"username" env:"SHORTY_SMTP_USERNAME"`
	Password string `json:"password" env:"SHORTY_SMTP_PASSWORD"`
	From     string `json:"from" env:"SHORTY_SMTP_FROM"`
}

type Tokens struct {
	Secret string        `json:"secret" env:"SHORTY_TOKEN_SECRET"`
	Expiry time.Duration `json:"expiry" env:"SHORTY_TOKEN_EXPIRY"`
}

type PsqlConfig struct {
	Host     string `json:"host" env:"SHORTY_DB_HOST" env-default:"localhost"`
	Port     uint16 `json:"port" env:"SHORTY_DB_PORT" env-default:"5432"`
	DB       string `json:"db" env:"SHORTY_DB" env-default:"shorty-local"`
	User     string `json:"user" env:"SHORTY_DB_USER" env-default:"shorty-local"`
	Password string `json:"password" env:"SHORTY_DB_PASSWORD" env-default:"shorty-local"`
}
