package common

import "errors"

var (
	ErrNotFound = errors.New("Not Found")
	ErrUnauthorized = errors.New("Unauthorized")
)