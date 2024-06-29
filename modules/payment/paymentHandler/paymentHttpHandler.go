package paymentHandler

import (
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandler interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecase
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecase) PaymentHttpHandler {
	return &paymentHttpHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
