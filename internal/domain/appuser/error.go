package appuser

import "github.com/sekalahita/epirus/internal/errors"

const (
	ErrorCodeUserNotFound    = "user-001"
	ErrorInvalidOnboardState = "user-002"
)

func NewErrorUserNotFound(err error) errors.CommonErr {
	return errors.CommonErr{
		Message:   "user not found",
		ErrorCode: ErrorCodeUserNotFound,
		Debug:     err,
	}
}

func NewErrorInvalidOnboardState(err error) errors.CommonErr {
	return errors.CommonErr{
		Message:   "invalid onboard state",
		ErrorCode: ErrorInvalidOnboardState,
		Debug:     err,
	}
}
