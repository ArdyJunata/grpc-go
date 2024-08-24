package internalerror

import (
	"errors"

	"google.golang.org/grpc/codes"
)

type Error struct {
	Code    codes.Code
	Message string
}

func (e Error) WithMessage(message string) Error {
	e.Message = message
	return e
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrUnauthenticated = errors.New("unauthorized")
	ErrNotFound        = errors.New("not found")
	ErrInternal        = errors.New("internal error")
	ErrAlreadyExists   = errors.New("already exists")
	ErrUnknown         = errors.New("unknown error")
)

var (
	ErrorUnauthenticated = Error{
		Code:    codes.Unauthenticated,
		Message: ErrUnauthenticated.Error(),
	}

	ErrorNotFound = Error{
		Code:    codes.NotFound,
		Message: ErrNotFound.Error(),
	}

	ErrorInternal = Error{
		Code:    codes.Internal,
		Message: ErrInternal.Error(),
	}

	ErrorAlreadyExists = Error{
		Code:    codes.AlreadyExists,
		Message: ErrAlreadyExists.Error(),
	}

	ErrorUnknown = Error{
		Code:    codes.Unknown,
		Message: ErrUnknown.Error(),
	}
)

var (
	ErrorBase = map[string]Error{
		ErrorUnauthenticated.Error(): ErrorUnauthenticated,
		ErrorNotFound.Error():        ErrorNotFound,
		ErrorInternal.Error():        ErrorInternal,
		ErrorUnknown.Error():         ErrorUnknown,
		ErrorAlreadyExists.Error():   ErrorAlreadyExists,
	}
)
