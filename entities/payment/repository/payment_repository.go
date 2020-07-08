package repository

import (
	"time"

	baseRepo "uit_payment/entities/base_repository"
	"uit_payment/enum"
	"uit_payment/model"
)

type PaymentRepository struct {
	baseRepo.BaseRepository
}

func NewPaymentRepository() PaymentRepositoryInterface {
	repo := &PaymentRepository{}
	repo.BaseRepository.Init()
	return repo
}

func (repo *PaymentRepository) FindByTransactionID(transactionID string, obj *model.Payment) error {
	return repo.BaseRepository.DB.
		Where(model.Payment{TransactionID: transactionID}).Last(obj).Error
}

func (repo *PaymentRepository) FindByPaymentTX(PaymentTX string, obj *model.Payment) error {
	return repo.BaseRepository.FindBy(model.Payment{PaymentTX: PaymentTX}, obj)
}

func (repo *PaymentRepository) Create(obj *model.Payment) error {
	return repo.BaseRepository.Create(obj)
}

func (repo *PaymentRepository) Update(obj *model.Payment) error {
	return repo.BaseRepository.Update(obj)
}
func (repo *PaymentRepository) CreateWithPaymentRequest(obj *model.Payment, paymentRequest *model.PaymentRequest) error {
	tx := repo.BaseRepository.DB.Begin()

	if err := tx.Create(obj).Error; err != nil {
		tx.Rollback()
		return err
	}

	paymentRequest.PaymentID = obj.ID

	if err := tx.Create(paymentRequest).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (repo *PaymentRepository) UpdatePaid(obj *model.Payment, paymentRequest *model.PaymentRequest) error {
	if obj.PaidAt == nil {
		t := time.Now()
		obj.PaidAt = &t
	}

	obj.Status = enum.PaymentStatusPaid

	tx := repo.BaseRepository.DB.Begin()
	if err := tx.Model(obj).Update(obj).Error; err != nil {

		tx.Rollback()

		return err
	}

	paymentRequest.PaymentID = obj.ID
	if err := tx.Create(paymentRequest).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (repo *PaymentRepository) UpdateFailed(obj *model.Payment, paymentRequest *model.PaymentRequest) error {
	obj.Status = enum.PaymentStatusFailed

	tx := repo.BaseRepository.DB.Begin()

	if err := tx.Model(obj).Update(obj).Error; err != nil {
		tx.Rollback()
		return err
	}

	paymentRequest.PaymentID = obj.ID
	if err := tx.Create(paymentRequest).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (repo *PaymentRepository) UpdateRefunded(obj *model.Payment, paymentRequest *model.PaymentRequest) error {
	obj.Status = enum.PaymentStatusRefund

	tx := repo.BaseRepository.DB.Begin()
	if err := tx.Model(obj).Update(obj).Error; err != nil {
		tx.Rollback()
		return err
	}

	paymentRequest.PaymentID = obj.ID
	if err := tx.Create(paymentRequest).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
