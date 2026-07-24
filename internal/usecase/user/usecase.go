package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	passwordDomain "github.com/lunadiotic/shopex-go/internal/domain/password"
	domain "github.com/lunadiotic/shopex-go/internal/domain/user"
)

type UseCase struct {
	repository domain.Repository
	hasher     passwordDomain.Hasher
}

func NewUseCase(
	repository domain.Repository,
	hasher passwordDomain.Hasher,
) *UseCase {
	return &UseCase{repository: repository, hasher: hasher}
}

func (uc *UseCase) Register(ctx context.Context, user *domain.User) error {
	if err := uc.validate(user); err != nil {
		return err
	}

	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	existingUser, err := uc.repository.FindByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existingUser != nil {
		return ErrEmailAlreadyExists
	}

	user.Password, err = uc.hasher.Hash(user.Password)
	if err != nil {
		return err
	}

	return uc.repository.Create(ctx, user)
}