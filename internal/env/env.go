package env

type ENV string

const (
	LOCAL ENV = "LOCAL"
	DEV   ENV = "DEV"
	PROD  ENV = "PROD"
)
