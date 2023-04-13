package models

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sekalahita/epirus/internal/errors"
)

func MarshalDate(dateTime time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if dateTime.IsZero() {
			w.Write([]byte(`""`))
			return
		}

		w.Write([]byte(strconv.Quote(dateTime.Format(time.DateOnly))))
	})
}

func UnmarshalDate(v interface{}) (time.Time, error) {
	vi, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("value is not a string: %T", v)
	}

	value, err := time.Parse(time.DateOnly, vi)
	return value, errors.ErrorWithCurrentFuncName(err)
}
