package storages

import (
	"context"
	"shorty/internal/core/domain/models"
)

type UserStorage interface {
	Add(ctx context.Context, args *AddUserArgs) error
	FindUsername(ctx context.Context, username string) (string, error)
	FindEmail(ctx context.Context, email string) (string, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	ActivateUser(ctx context.Context, url string) error
}

type AddUserArgs struct {
	ID             string
	Email          string
	Username       string
	Password       string
	ActivationCode string
}
