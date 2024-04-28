package store

import "errors"

var (
	// ErrorRecordNotFound ...
	ErrorRecordNotFound = errors.New("record(s) not found")
	// ErrorNotExistRef ...
	ErrorNotExistRef = errors.New("given reference doesn't exist")
	// ErrorNoAccess ...
	ErrorNoAccess = errors.New("insufficient user rights")
)
