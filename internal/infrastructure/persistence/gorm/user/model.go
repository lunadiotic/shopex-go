package user

import "time"

type Model struct {
	ID uint `gorm:"primaryKey"`

	FullName string
	Email string
	Password string

	Status string
	
	CreatedAt time.Time
	UpdatedAt time.Time
}