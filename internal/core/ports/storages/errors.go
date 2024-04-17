package storages

import "errors"

var (
	ErrNotFound       = errors.New("storage: not found")
	ErrInternal       = errors.New("storage: internal error")
	ErrDuplicateValue = errors.New("storage: duplicate value")
)
