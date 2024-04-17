package storages

import "context"

type ApiKeyStorage interface {
	Add(context.Context, *Key) error
	Delete(context.Context) error
}

type Key struct {
	id  string
	key string
}
