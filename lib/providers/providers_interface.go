package providers

import (
	"uit_payment/entities/payment/delivery/request"
	"uit_payment/model"
)

// Provider - Payment Provider
type ProviderInterface interface {
	Name() string
	CreatePayment(paymentRequest *request.CreatePaymentRequest, paymentModel *model.Payment) (*model.PaymentRequest, error)
	GetPayment(paymentModel *model.Payment) (*model.PaymentRequest, error)
	RefundPayment(paymentModel *model.Payment) (*model.PaymentRequest, error)
}

type Provider interface {
	GetProvider(mpayment *model.Payment) ProviderInterface
}
