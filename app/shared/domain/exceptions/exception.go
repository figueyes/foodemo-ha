package exceptions

import "errors"

var (
	ErrorConnectionDB = errors.New("error trying to save in database")
)