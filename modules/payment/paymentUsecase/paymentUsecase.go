package paymentUsecase

import "github.com/Montheankul-K/game-microservices/modules/payment/paymentRepository"

type (
	PaymentUsecase interface{}

	paymentUsecase struct {
		paymentRepository paymentRepository.PaymentRepository
	}
)

func NewPaymentUsecase(paymentRepository paymentRepository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
	}
}
