package config

import "time"

type HttpServer struct {
	Host         string        `toml:"host"`
	Port         int           `toml:"port"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
	IdleTimeout  time.Duration `toml:"idle_timeout"`
}

func HttpServerFromConfig(cfg *Config) *HttpServer { return cfg.HttpServer }
