package repository

import "uit_payment/model"

type PartnerRepositoryInterface interface {
	FindByKey(key string, obj *model.Partner) error
}
