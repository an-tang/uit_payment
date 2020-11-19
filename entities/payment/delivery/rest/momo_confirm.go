package http

import (
	"context"
	"net/http"
	"strconv"

	partnerRepo "uit_payment/entities/partner"
	"uit_payment/entities/payment/delivery/request"
	repository "uit_payment/entities/payment/repository"
	service "uit_payment/entities/payment/service"
	paymentRequestRepo "uit_payment/entities/payment_request/repository"
	"uit_payment/enum"
	"uit_payment/lib/env"
	"uit_payment/lib/hmac"
	httpclient "uit_payment/lib/http_client"
	"uit_payment/lib/logging"
	log "uit_payment/lib/logging"
	"uit_payment/lib/providers/momo"
	"uit_payment/model"
	dom_api "uit_payment/services/dom"
)

type MomoConfirm struct {
	Handler
	PaymentService     service.PaymentServiceInterface
	PaymentRepo        repository.PaymentRepositoryInterface
	PaymentRequestRepo paymentRequestRepo.PaymentRequestRepositoryInterface
	PartnerRepo        partnerRepo.PartnerRepositoryInterface
	DOMService         dom_api.DOMServiceInterface
	HTTPClient         httpclient.HTTPCInterface
}

func NewMomoConfirm() *MomoConfirm {
	return &MomoConfirm{
		PaymentRepo:        repository.NewPaymentRepository(),
		PaymentRequestRepo: paymentRequestRepo.NewPaymentRequestRepository(),
		PaymentService:     service.NewPaymentService(),
		HTTPClient:         httpclient.HTTPInstance(),
		DOMService:         dom_api.NewDOMService(),
	}
}

func (m *MomoConfirm) Handle() {
	params := &momo.MomoConfirmPaymentRequest{}
	m.ParseParam(&params)

	payment := &model.Payment{}
	err := m.PaymentRepo.FindByPaymentTX(params.RequestID, payment)
	if err != nil {
		log.WithError(err).Errorln("MomoConfirm.CannotFindTransaction: " + params.RequestID)
		return
	}

	paymentRequestLog := &model.PaymentRequest{
		RequestType: enum.PaymentRequestTypeWebhook,
		PaymentID:   payment.ID,
	}

	resp := parseConfirm(*params)
	if resp.ErrorCode != 0 {
		paymentRequestLog.Populate(params, resp, http.StatusBadRequest)
		m.PaymentRequestRepo.Create(paymentRequestLog)
		m.RenderErrorWithJSON(resp)
		return
	}

	m.RenderSuccess(resp)

	paymentRequestLog.Populate(params, resp, http.StatusOK)
	// go m.callbackPartner(params, payment)
	go m.callbackgRPC(*payment)

	go m.updatePayment(payment, paymentRequestLog, params)
}

func parseConfirm(req momo.MomoConfirmPaymentRequest) momo.MomoAIOConfirmResponse {
	resp := momo.MomoAIOConfirmResponse{
		AccessKey:    req.AccessKey,
		PartnerCode:  req.PartnerCode,
		OrderID:      req.OrderID,
		RequestID:    req.RequestID,
		ResponseTime: req.ResponseTime,
		ErrorCode:    req.ErrorCode,
		Message:      req.Message,
	}

	hmacData := resp.HmacCombine()
	resp.Signature = hmac.HexStringEncode(hmac.SHA256, env.GetMomoSecretKey(), hmacData)
	return resp
}

func (m *MomoConfirm) updatePayment(payment *model.Payment, paymentRequestLog *model.PaymentRequest, params *momo.MomoConfirmPaymentRequest) {
	if params.ErrorCode != 0 {
		m.PaymentService.UpdateFailed(payment, paymentRequestLog)
	} else {
		m.PaymentService.UpdatePaid(payment, paymentRequestLog)
	}
}

func (m *MomoConfirm) callbackgRPC(payment model.Payment) {
	payment.Status = enum.PaymentStatusPaid
	m.DOMService.PaymentCallback(context.TODO(), payment)
}

func (m *MomoConfirm) callbackPartner(params *momo.MomoConfirmPaymentRequest, payment *model.Payment) {
	amount, err := strconv.ParseFloat(params.Amount, 32)
	if err != nil {
		logging.WithError(err).Errorln("MomoConfirm.ConvertAmountError:", err.Error())
		amount = 0
	}

	req := request.CallbackPartnerRequest{
		TransactionID: payment.TransactionID,
		Amount:        float32(amount),
		Status:        enum.PaymentStatusPaid,
	}

	if params.ErrorCode != 0 {
		req.Status = enum.PaymentStatusFailed
	}

	partner, err := m.PartnerRepo.FindByID(payment.PartnerID)
	if err != nil {
		logging.WithError(err).Errorln("FindPartnerError:", err.Error())
	}

	endpoint := partner.CallbackURL
	err = m.HTTPClient.Post(endpoint, "application/json", req, nil)
	if err != nil {
		logging.WithError(err).Errorln("MomoConfirm.CallbackPartnerError:", err.Error())
	}
}
