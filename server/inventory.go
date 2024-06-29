package server

import (
	"github.com/Montheankul-K/game-microservices/modules/inventory/inventoryHandler"
	"github.com/Montheankul-K/game-microservices/modules/inventory/inventoryRepository"
	"github.com/Montheankul-K/game-microservices/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repository := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repository)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")
	inventory.GET("", s.healthCheckService)
}
