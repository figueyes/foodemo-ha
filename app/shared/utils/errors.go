package utils

import (
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

func ErrorValid(err error) error {
	var msgError string
	var split string
	for _, e := range err.(validator.ValidationErrors) {
		msgError = fmt.Sprintf("%s%s%s", msgError, split, e)
		split = ","
	}
	return errors.New(msgError)
}
