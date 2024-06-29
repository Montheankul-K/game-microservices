package middlewareHandler

import (
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandler interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecase
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecase) MiddlewareHandler {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}
