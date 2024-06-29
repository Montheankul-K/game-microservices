package server

import (
	"github.com/Montheankul-K/game-microservices/modules/item/itemHandler"
	"github.com/Montheankul-K/game-microservices/modules/item/itemRepository"
	"github.com/Montheankul-K/game-microservices/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repository := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repository)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")
	item.GET("", s.healthCheckService)
}
