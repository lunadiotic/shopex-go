package user

import (
	"context"

	domain "github.com/lunadiotic/shopex-go/internal/domain/user"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB	
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

var _ domain.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, user *domain.User) error {
	panic("not implemented")
}

func (r *Repository) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	panic("not implemented")
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("not implemented")
}

func (r *Repository) Update(ctx context.Context, user *domain.User) error {
	panic("not implemented")
}

func (r *Repository) Delete(ctx context.Context, id uint) error {
	panic("not implemented")
}