package config

type RefreshTokensStorage struct {
	DSN      string `toml:"dsn"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

func RefreshTokensStorageFromConfig(cfg *Config) *RefreshTokensStorage {
	return cfg.RefreshTokensStorage
}
