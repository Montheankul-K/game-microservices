package server

import (
	"github.com/Montheankul-K/game-microservices/modules/inventory/inventoryHandler"
	inventoryPb "github.com/Montheankul-K/game-microservices/modules/inventory/inventoryPb"
	"github.com/Montheankul-K/game-microservices/modules/inventory/inventoryRepository"
	"github.com/Montheankul-K/game-microservices/modules/inventory/inventoryUsecase"
	"github.com/Montheankul-K/game-microservices/pkg/grpccon"
	"log"
)

func (s *server) inventoryService() {
	repository := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repository)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	go func() {
		grpcServer, listen := grpccon.NewGrpcServer(s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)
		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("inventory gRPC server is listening on %s", s.cfg.Grpc.InventoryUrl)
		grpcServer.Serve(listen)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")
	inventory.GET("", s.healthCheckService)
}
