package store

import "errors"

var (
	// ErrorRecordNotFound ...
	ErrorRecordNotFound = errors.New("record not found")
	ErrorNotExistRef    = errors.New("given reference doesn't exist")
)
