package server

import (
	"github.com/Montheankul-K/game-microservices/modules/auth/authHandler"
	authPb "github.com/Montheankul-K/game-microservices/modules/auth/authPb"
	"github.com/Montheankul-K/game-microservices/modules/auth/authRepository"
	"github.com/Montheankul-K/game-microservices/modules/auth/authUsecase"
	"github.com/Montheankul-K/game-microservices/pkg/grpccon"
	"log"
)

func (s *server) authService() {
	repository := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repository)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	go func() {
		grpcServer, listen := grpccon.NewGrpcServer(s.cfg.Jwt, s.cfg.Grpc.AuthUrl)
		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("auth gRPC server is listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(listen)
	}()

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")
	auth.GET("", s.healthCheckService)
}
