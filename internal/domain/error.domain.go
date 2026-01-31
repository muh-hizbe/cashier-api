package domain

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrInternal       = errors.New("internal error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrDuplicate      = errors.New("duplicate")
	ErrConflict       = errors.New("conflict")
	ErrNotImplemented = errors.New("not implemented")
)
