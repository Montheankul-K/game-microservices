package playerHandler

import (
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/modules/player/playerUsecase"
)

type (
	PlayerHttpHandler interface{}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecase
	}
)

func NewPlayerHttpHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecase) PlayerHttpHandler {
	return &playerHttpHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}
