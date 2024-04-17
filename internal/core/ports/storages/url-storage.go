package storages

import (
	"context"
	"shorty/internal/core/domain/models"
)

type UrlStorage interface {
	Add(context.Context, *AddUrlArgs) error
	Delete(context.Context, *UrlOptions) error
	ReplaceLong(context.Context, *ReplaceLongUrlArgs) error
	FindOne(context.Context, *UrlOptions) (*models.URL, error)
	FindOwned(context.Context, string) ([]models.URL, error)
}

type AddUrlArgs struct {
	ID    string
	Owner string
	Short string
	Long  string
}

type UrlOptions struct {
	Long  *string
	Short *string
}

type ReplaceLongUrlArgs struct {
	To   string
	From string
}
