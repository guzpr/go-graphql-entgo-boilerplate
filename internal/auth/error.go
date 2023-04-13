package auth

import "github.com/sekalahita/epirus/internal/errors"

const (
	ErrorCodeAuthTokenNotFound = "auth-001"
)

func NewErrorAuthTokenNotFound(err error) errors.CommonErr {
	return errors.CommonErr{
		Message:   "auth token not found",
		ErrorCode: ErrorCodeAuthTokenNotFound,
		Debug:     err,
	}
}
