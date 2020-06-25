package repository

import (
	baseRepo "uit_payment/entities/base_repository"
	"uit_payment/model"
)

type PaymentRequestRepository struct {
	baseRepo.BaseRepository
}

func NewPaymentRequestRepository() *PaymentRequestRepository {
	repo := &PaymentRequestRepository{}
	repo.Init()
	return repo
}

func (repo *PaymentRequestRepository) Create(obj *model.PaymentRequest) error {
	return repo.BaseRepository.Create(obj)
}
