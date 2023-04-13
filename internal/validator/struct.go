package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sekalahita/epirus/internal/errors"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func ValidateStruct(s any) error {
	if err := v.Struct(s); err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			verrs := err.(validator.ValidationErrors)
			errs := make([]string, len(verrs))

			for i, verr := range verrs {
				// TODO: do a reverse lookup of the json key
				errs[i] = fmt.Sprintf(
					"%s: failed on '%s' tag.",
					verr.StructNamespace(),
					verr.ActualTag(),
				)
			}

			return errors.ErrorWithCurrentFuncName(errors.New(strings.Join(errs, " ")))
		default:
			return errors.ErrorWithCurrentFuncName(err)
		}
	}

	return nil
}
