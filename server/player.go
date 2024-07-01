package server

import (
	"github.com/Montheankul-K/game-microservices/modules/player/playerHandler"
	playerPb "github.com/Montheankul-K/game-microservices/modules/player/playerPb"
	"github.com/Montheankul-K/game-microservices/modules/player/playerRepository"
	"github.com/Montheankul-K/game-microservices/modules/player/playerUsecase"
	"github.com/Montheankul-K/game-microservices/pkg/grpccon"
	"log"
)

func (s *server) playerService() {
	repository := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repository)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)

	go func() {
		grpcServer, listen := grpccon.NewGrpcServer(s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)
		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("player gRPC server is listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(listen)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")
	player.GET("", s.healthCheckService)
}
