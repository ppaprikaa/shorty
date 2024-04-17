package storages

import "context"

type TokenStorage interface {
	FindSession(ctx context.Context, refreshToken string) (string, error)
	Delete(context.Context, *DeleteTokenArgs) error
	Add(context.Context, *AddTokenArgs) error
}

type DeleteTokenArgs struct {
	SessionID string
}

type AddTokenArgs struct {
	ID        string
	Token     string
	SessionID string
}
