package apiserver

import "errors"

var (
	ErrorIncorrectEmailOrPassword = errors.New("incorrect email or password")
	ErrorNotAuthenticated         = errors.New("not authenticated")
)
