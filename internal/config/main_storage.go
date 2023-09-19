package config

type MainStorage struct {
	DSN string `toml:"dsn"`
}

func MainStorageFromConfig(cfg *Config) *MainStorage {
	return cfg.MainStorage
}
