package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrInputBinding        = errors.New("wrong parameter")
	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("bad request")
	ErrForbidden           = errors.New("forbidden request")
)
