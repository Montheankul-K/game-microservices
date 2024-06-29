package authUsecase

import "github.com/Montheankul-K/game-microservices/modules/auth/authRepository"

type (
	AuthUsecase interface{}

	authUsecase struct {
		authRepository authRepository.AuthRepository
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepository) AuthUsecase {
	return &authUsecase{
		authRepository: authRepository,
	}
}
