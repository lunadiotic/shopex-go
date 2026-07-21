package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *Hasher) Compare(hashedPassword, password string) (bool, error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil, nil
}