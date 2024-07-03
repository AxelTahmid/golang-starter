package message

import "errors"

var (
	ErrBadRequest    = errors.New("error bad request")
	ErrInternalError = errors.New("internal server error")

	ErrFormingResponse = errors.New("error forming response")

	ErrNoRecord = errors.New("no record found")

	ErrPassOrUserIncorrect = errors.New("password or username is incorrect")
)
