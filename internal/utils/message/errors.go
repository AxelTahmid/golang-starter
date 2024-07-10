package message

import "errors"

// errors
var (
	ErrBadRequest          = errors.New("error bad request")
	ErrInternalError       = errors.New("internal server error")
	ErrFormingResponse     = errors.New("error forming response")
	ErrUnauthorized        = errors.New("unauthorized access")
	ErrNoRecord            = errors.New("no record found")
	ErrPassOrUserIncorrect = errors.New("password or email is incorrect")
	ErrBadTokenFormat      = errors.New("authorization header format must be 'Bearer {token}'")
)
