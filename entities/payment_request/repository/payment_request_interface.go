package repository

import "uit_payment/model"

type PaymentRequestRepositoryInterface interface {
	Create(obj *model.PaymentRequest) error
}
