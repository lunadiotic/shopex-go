package user

import (
	"net/mail"
	"strings"

	domain "github.com/lunadiotic/shopex-go/internal/domain/user"
)

func (uc *UseCase) validate(user *domain.User) error {
	if strings.TrimSpace(user.FullName) == "" {
		return ErrFullNameRequired
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return ErrInvalidEmail
	}

	if strings.TrimSpace(user.Password) == "" {
		return ErrPasswordRequired
	}

	if len(user.Password) < 8 {
		return ErrPasswordTooShort
	}

	return nil
 }