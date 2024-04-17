package models

type User struct {
	ID             string
	Email          string
	Username       string
	Password       string
	ActivationCode string
	Activated      bool
}
