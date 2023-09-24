package config

type Mailer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

func MailerFromConfig(cfg *Config) *Mailer {
	return cfg.Mailer
}
