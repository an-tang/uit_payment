package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"uit_payment/entities/payment/delivery/request"
	repository "uit_payment/entities/payment/repository"
	paymentRequestRepo "uit_payment/entities/payment_request/repository"

	"uit_payment/enum"
	"uit_payment/lib/providers"
	"uit_payment/lib/providers/provider"
	"uit_payment/model"
)

type PaymentService struct {
	PaymentClient      providers.ProviderInterface
	Provider           providers.Provider
	PaymentRepo        repository.PaymentRepositoryInterface
	PaymentRequestRepo paymentRequestRepo.PaymentRequestRepositoryInterface
}

// NewpaymentUsecase :
func NewPaymentService() PaymentServiceInterface {
	return &PaymentService{
		PaymentRepo:        repository.NewPaymentRepository(),
		PaymentRequestRepo: paymentRequestRepo.NewPaymentRequestRepository(),
		Provider:           provider.NewProvider(),
	}
}

func (p *PaymentService) CreatePayment(mpaymentRequest *request.CreatePaymentRequest, mpayment *model.Payment) (*model.Payment, error) {
	mpayment = p.popuplateModel(mpaymentRequest)
	err := p.PaymentRepo.FindByTransactionID(mpayment.TransactionID, mpayment)
	if err == nil {
		return mpayment, fmt.Errorf("TransactionID Exist: %s", mpayment.TransactionID)
	}

	p.PaymentClient = p.Provider.GetProvider(mpayment)
	paymentRequestLog, err := p.PaymentClient.CreatePayment(mpaymentRequest, mpayment)
	if err != nil {
		log.Printf("Create payment Error : %v", err)
		return mpayment, err
	}

	mpayment.UID = mpayment.GenerateUID()
	err = p.PaymentRepo.CreateWithPaymentRequest(mpayment, paymentRequestLog)
	if err != nil {
		log.Printf("Create payment Error : %v", err)
		return mpayment, err
	}

	log.Printf("Create payment: %v", mpayment)
	return mpayment, nil
}

func (p *PaymentService) GetPayment(transactionID string) (*model.Payment, error) {
	mpayment := &model.Payment{}
	err := p.PaymentRepo.FindByTransactionID(transactionID, mpayment)
	if err != nil {
		return nil, err
	}
	if !mpayment.IsPersisted() {
		return nil, fmt.Errorf("Invalid %s", transactionID)
	}

	if mpayment.Status == enum.PaymentStatusPaid || mpayment.Status == enum.PaymentStatusRefund {
		return mpayment, nil
	}

	p.PaymentClient = p.Provider.GetProvider(mpayment)

	paymentRequestLog, err := p.PaymentClient.GetPayment(mpayment)

	if err != nil {
		log.Printf("Get payment Error : %v", err)
		return mpayment, err
	}

	if mpayment.Status == enum.PaymentStatusNew {
		p.PaymentRequestRepo.Create(paymentRequestLog)
		return mpayment, nil
	}

	if mpayment.Status == enum.PaymentStatusPaid {
		err = p.PaymentRepo.UpdatePaid(mpayment, paymentRequestLog)
		if err != nil {
			return nil, err
		}
		return mpayment, nil
	}

	if mpayment.Status == enum.PaymentStatusWaitingForPayment {
		err = p.PaymentRepo.UpdateWaitingForPayment(mpayment, paymentRequestLog)
		if err != nil {
			return nil, err
		}
		return mpayment, nil
	}

	err = p.PaymentRepo.UpdateFailed(mpayment, paymentRequestLog)
	if err != nil {
		return mpayment, err
	}

	return mpayment, nil
}

func (p *PaymentService) RefundPayment(transactionID string) (*model.Payment, error) {
	mpayment := &model.Payment{}
	err := p.PaymentRepo.FindByTransactionID(transactionID, mpayment)
	if err != nil {
		return nil, err
	}

	if !mpayment.IsPersisted() {
		return nil, fmt.Errorf("Invalid %s", transactionID)
	}

	p.PaymentClient = p.Provider.GetProvider(mpayment)
	paymentRequestLog, err := p.PaymentClient.RefundPayment(mpayment)
	if err != nil {
		err2 := p.PaymentRequestRepo.Create(paymentRequestLog)
		if err2 != nil {
			return mpayment, err2
		}
		return mpayment, err

	}

	err = p.PaymentRepo.UpdateRefunded(mpayment, paymentRequestLog)
	if err != nil {
		err2 := p.PaymentRequestRepo.Create(paymentRequestLog)
		log.Printf("====Payment===%v", err2)
		if err2 != nil {
			return mpayment, err2
		}
		return mpayment, err

	}
	return mpayment, nil
}

func (p *PaymentService) popuplateModel(param *request.CreatePaymentRequest) *model.Payment {
	return &model.Payment{
		Currency:      param.Currency,
		TransactionID: param.TransactionID,
		PaymentMethod: param.PaymentMethod,
		Amount:        param.Amount,
		StoreID:       param.StoreID,
		Status:        1,
	}
}

func (p *PaymentService) UpdatePaid(payment *model.Payment, paymentRequest *model.PaymentRequest) error {
	if !payment.IsPersisted() {
		paymentRequest.Status = http.StatusBadRequest
		p.PaymentRequestRepo.Create(paymentRequest)
		return errors.New("Payment is persisted")
	}

	return p.PaymentRepo.UpdatePaid(payment, paymentRequest)
}

func (p *PaymentService) UpdateFailed(payment *model.Payment, paymentRequest *model.PaymentRequest) error {
	if !payment.IsPersisted() {
		paymentRequest.Status = http.StatusBadRequest
		p.PaymentRequestRepo.Create(paymentRequest)
		return errors.New("Payment is persisted")
	}

	return p.PaymentRepo.UpdateFailed(payment, paymentRequest)
}
