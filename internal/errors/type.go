package errors

import "fmt"

type CommonErr struct {
	Message   string
	Debug     error
	ErrorCode string
}

func (e CommonErr) Error() string {
	return fmt.Sprintf("code: %s,message: %s, debug: %s", e.ErrorCode, e.Message, e.Debug.Error())
}

type HttpErr struct {
	CommonErr
	HttpCode int
}

func (e HttpErr) Error() string {
	return fmt.Sprintf("http code: %d, error code: %s, message: %s, debug: %s", e.HttpCode, e.ErrorCode, e.Message, e.Debug.Error())
}
