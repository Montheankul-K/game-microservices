package server

import (
	"github.com/Montheankul-K/game-microservices/modules/item/itemHandler"
	itemPb "github.com/Montheankul-K/game-microservices/modules/item/itemPb"
	"github.com/Montheankul-K/game-microservices/modules/item/itemRepository"
	"github.com/Montheankul-K/game-microservices/modules/item/itemUsecase"
	"github.com/Montheankul-K/game-microservices/pkg/grpccon"
	"log"
)

func (s *server) itemService() {
	repository := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repository)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	go func() {
		grpcServer, listen := grpccon.NewGrpcServer(s.cfg.Jwt, s.cfg.Grpc.ItemUrl)
		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("item gRPC server is listening on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(listen)
	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")
	item.GET("", s.healthCheckService)
}
