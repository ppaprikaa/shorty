package env

type ENV string

// types of ENV
const (
	LOCAL ENV = "LOCAL"
	DEV   ENV = "DEV"
	PROD  ENV = "PROD"
)
