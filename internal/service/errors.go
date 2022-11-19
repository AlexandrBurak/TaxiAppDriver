package service

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrWrongPassword        = errors.New("wrong password")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrContextMissingDriver = errors.New("no driver in context")
)
