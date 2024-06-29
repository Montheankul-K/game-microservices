package playerHandler

import (
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/modules/player/playerUsecase"
)

type (
	PlayerQueueHandler interface{}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecase
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecase) PlayerQueueHandler {
	return &playerQueueHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}
