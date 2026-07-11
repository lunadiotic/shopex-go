package user

import domain "github.com/lunadiotic/shopex-go/internal/domain/user"

type UseCase struct {
	repository domain.Repository
}

func NewUseCase(repository domain.Repository) *UseCase {
	return &UseCase{repository: repository}
}

func (uc *UseCase) Ping() string {
	return "user usecase is ready"
}