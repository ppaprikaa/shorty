package tokenizer

import (
	"context"
	"errors"
	"shorty/internal/core/domain/models"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrInternal     = errors.New("internal error")
)

type Tokenizer interface {
	GenerateAccess(ctx context.Context, claims *models.AccessTokenClaims) (*models.AccessToken, error)
	GenerateRefresh(ctx context.Context, claims *models.RefreshTokenClaims) (*models.RefreshToken, error)
	VerifyAccess(ctx context.Context, token string) (*models.AccessTokenClaims, error)
	VerifyRefresh(ctx context.Context, token string) (*models.RefreshTokenClaims, error)
}
