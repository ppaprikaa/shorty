package user

import (
	"context"
	"errors"
)

var (
	ErrTakenUsernameAndEmail = errors.New("username and email are already taken")
	ErrTakenUsername         = errors.New("username already taken")
	ErrTakenEmail            = errors.New("email already taken")
)

type Service interface {
	Registrate(ctx context.Context, data RegistrationDTO) error
}
