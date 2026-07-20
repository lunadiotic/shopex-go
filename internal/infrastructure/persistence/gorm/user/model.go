package user

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	FullName string `gorm:"size:100;not null"`
	Email    string `gorm:"size:255;uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Status   string `gorm:"size:20;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Model) TableName() string {
	return "users"
}