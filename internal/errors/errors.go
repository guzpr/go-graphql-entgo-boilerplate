package errors

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

const (
	ErrorCodeInternalServer = "srv-001"
)

func NewErrorInternalServer(err error) CommonErr {
	return CommonErr{
		Message:   "internal server error",
		ErrorCode: ErrorCodeInternalServer,
		Debug:     err,
	}
}

func ErrorWithCurrentFuncName(err error) error {
	return errors.Wrap(err, FuncNameFromFrames(3))
}

func FuncNameFromFrames(skip int) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	str := strings.Split(frame.Func.Name(), "/")
	return str[len(str)-1]
}

func IsNotFoundError(err error) bool {
	return true
}

func New(message string) error {
	return errors.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Wrap(err error, message string) error {
	if err == nil {
		log.Printf("passing nil error for: %s", message)
		return errors.New(message)
	}

	return errors.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		log.Printf("passing nil error for: %s", fmt.Sprintf(format, args...))
		return errors.Errorf(format, args...)
	}

	return errors.Wrapf(err, format, args...)
}
