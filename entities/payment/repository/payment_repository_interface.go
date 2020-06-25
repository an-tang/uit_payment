package repository

import "uit_payment/model"

type PaymentRepositoryInterface interface {
	Create(obj *model.Payment) error
	CreateWithPaymentRequest(obj *model.Payment, paymentRequest *model.PaymentRequest) error
	UpdatePaid(obj *model.Payment, paymentRequest *model.PaymentRequest) error
	UpdateFailed(obj *model.Payment, paymentRequest *model.PaymentRequest) error
	FindByTransactionID(transactionID string, obj *model.Payment) error
	FindByPaymentTX(PaymentTX string, obj *model.Payment) error
	UpdateRefunded(obj *model.Payment, paymentRequest *model.PaymentRequest) error
	UpdateWaitingForPayment(obj *model.Payment, paymentRequest *model.PaymentRequest) error
	Update(obj *model.Payment) error
}
