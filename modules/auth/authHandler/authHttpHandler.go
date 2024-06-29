package authHandler

import (
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/modules/auth/authUsecase"
)

type (
	AuthHttpHandler interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecase
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecase) AuthHttpHandler {
	return &authHttpHandler{
		cfg:         cfg,
		authUsecase: authUsecase,
	}
}
