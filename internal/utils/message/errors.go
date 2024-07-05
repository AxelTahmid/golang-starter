package message

import "errors"

// errors
var (
	ErrBadRequest    = errors.New("error bad request")
	ErrInternalError = errors.New("internal server error")

	ErrFormingResponse = errors.New("error forming response")

	ErrNoRecord = errors.New("no record found")

	ErrPassOrUserIncorrect = errors.New("password or email is incorrect")
)

// success messages
var (
	SuccessLogin    = "user login successful ->"
	SuccessRegister = "user register successful ->"
)
