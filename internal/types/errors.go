package types

import "errors"

type StatusError struct {
	error
	Status int
}

func NewStatusError(msg string, status int) StatusError {
	return StatusError{
		error:  errors.New(msg),
		Status: status,
	}
}
