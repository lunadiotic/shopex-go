package user

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrFullNameRequired = errors.New("full name is required")
	ErrInvalidEmail = errors.New("invalid email")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
)