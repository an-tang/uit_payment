package service

import (
	"errors"
	"fmt"
	"net/http"

	partnerRepo "uit_payment/entities/partner"
	"uit_payment/entities/payment/delivery/request"
	repository "uit_payment/entities/payment/repository"
	paymentRequestRepo "uit_payment/entities/payment_request/repository"
	"uit_payment/enum"
	"uit_payment/lib/logging"
	"uit_payment/lib/providers"
	"uit_payment/lib/providers/provider"
	"uit_payment/model"
)

type PaymentService struct {
	PaymentClient      providers.ProviderInterface
	Provider           providers.Provider
	PaymentRepo        repository.PaymentRepositoryInterface
	PaymentRequestRepo paymentRequestRepo.PaymentRequestRepositoryInterface
	PartnerRepo        partnerRepo.PartnerRepositoryInterface
}

// NewpaymentUsecase :
func NewPaymentService() PaymentServiceInterface {
	return &PaymentService{
		PaymentRepo:        repository.NewPaymentRepository(),
		PaymentRequestRepo: paymentRequestRepo.NewPaymentRequestRepository(),
		PartnerRepo:        partnerRepo.NewPartnerRepository(),
		Provider:           provider.NewProvider(),
	}
}

func (p *PaymentService) CreatePayment(paymentRequest *request.CreatePaymentRequest, payment *model.Payment) (*model.Payment, error) {
	payment = p.popuplateModel(paymentRequest)
	if payment.PartnerID == 0 {
		return payment, fmt.Errorf("Cannot verify partner key: %s", paymentRequest.PartnerKey)
	}

	err := p.PaymentRepo.FindByTransactionID(payment.TransactionID, payment)
	if err == nil {
		return payment, fmt.Errorf("TransactionID Exist: %s", payment.TransactionID)
	}

	p.PaymentClient = p.Provider.GetProvider(payment)
	paymentRequestLog, err := p.PaymentClient.CreatePayment(paymentRequest, payment)
	if err != nil {
		return payment, err
	}

	payment.UID = payment.GenerateUID()
	err = p.PaymentRepo.CreateWithPaymentRequest(payment, paymentRequestLog)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (p *PaymentService) GetPayment(transactionID string) (*model.Payment, error) {
	payment := &model.Payment{}
	err := p.PaymentRepo.FindByTransactionID(transactionID, payment)
	if err != nil {
		return nil, err
	}

	if !payment.IsPersisted() {
		return nil, fmt.Errorf("Invalid %s", transactionID)
	}

	if payment.Status == enum.PaymentStatusPaid || payment.Status == enum.PaymentStatusRefund {
		return payment, nil
	}

	p.PaymentClient = p.Provider.GetProvider(payment)

	paymentRequestLog, err := p.PaymentClient.GetPayment(payment)
	if err != nil {
		return payment, err
	}

	if payment.Status == enum.PaymentStatusNew {
		p.PaymentRequestRepo.Create(paymentRequestLog)
		return payment, nil
	}

	if payment.Status == enum.PaymentStatusPaid {
		err = p.PaymentRepo.UpdatePaid(payment, paymentRequestLog)
		if err != nil {
			return nil, err
		}
		return payment, nil
	}

	err = p.PaymentRepo.UpdateFailed(payment, paymentRequestLog)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (p *PaymentService) RefundPayment(transactionID string) (*model.Payment, error) {
	payment := &model.Payment{}
	err := p.PaymentRepo.FindByTransactionID(transactionID, payment)
	if err != nil {
		return nil, err
	}

	if !payment.IsPersisted() {
		return nil, fmt.Errorf("Invalid %s", transactionID)
	}

	p.PaymentClient = p.Provider.GetProvider(payment)
	paymentRequestLog, err := p.PaymentClient.RefundPayment(payment)
	if err != nil {
		return payment, err
	}

	err = p.PaymentRepo.UpdateRefunded(payment, paymentRequestLog)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (p *PaymentService) popuplateModel(param *request.CreatePaymentRequest) *model.Payment {
	partner := p.findPartnerByKey(param.PartnerKey)

	return &model.Payment{
		Currency:      "VND",
		TransactionID: param.TransactionID,
		PaymentMethod: param.PaymentMethod,
		Amount:        param.Amount,
		StoreID:       param.StoreID,
		Status:        enum.PaymentStatusNew,
		PartnerID:     partner.ID,
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

func (p *PaymentService) findPartnerByKey(key string) model.Partner {
	partner := &model.Partner{}
	err := p.PartnerRepo.FindByKey(key, partner)
	if err != nil {
		logging.WithError(err).Errorln("PaymentService.findPartnerByKey:", err.Error())
		return model.Partner{}
	}

	return *partner
}
