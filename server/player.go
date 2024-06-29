package server

import (
	"github.com/Montheankul-K/game-microservices/modules/player/playerHandler"
	"github.com/Montheankul-K/game-microservices/modules/player/playerRepository"
	"github.com/Montheankul-K/game-microservices/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repository := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repository)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")
	player.GET("", s.healthCheckService)
}
