package server

import (
	"github.com/Montheankul-K/game-microservices/modules/payment/paymentHandler"
	"github.com/Montheankul-K/game-microservices/modules/payment/paymentRepository"
	"github.com/Montheankul-K/game-microservices/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repository := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repository)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")
	payment.GET("", s.healthCheckService)
}
