package service

import (
	"uit_payment/entities/payment/delivery/request"
	"uit_payment/model"
)

//PaymentServiceInterface interface payment service -
type PaymentServiceInterface interface {
	CreatePayment(mpaymentRequest *request.CreatePaymentRequest, mpayment *model.Payment) (*model.Payment, error)
	GetPayment(transactionID string) (*model.Payment, error)
	RefundPayment(transactionID string) (*model.Payment, error)
	UpdatePaid(obj *model.Payment, paymentRequest *model.PaymentRequest) error
	UpdateFailed(obj *model.Payment, paymentRequest *model.PaymentRequest) error
}
