package playerUsecase

import "github.com/Montheankul-K/game-microservices/modules/player/playerRepository"

type (
	PlayerUsecase interface{}

	playerUsecase struct {
		playerRepository playerRepository.PlayerRepository
	}
)

func NewPlayerUsecase(playerRepository playerRepository.PlayerRepository) PlayerUsecase {
	return &playerUsecase{
		playerRepository: playerRepository,
	}
}
