package middlewareUsecase

import "github.com/Montheankul-K/game-microservices/modules/middleware/middlewareRepository"

type (
	MiddlewareUsecase interface{}

	middlewareUsecase struct {
		middlewareRepository middlewareRepository.MiddlewareRepository
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepository) MiddlewareUsecase {
	return &middlewareUsecase{
		middlewareRepository: middlewareRepository,
	}
}
