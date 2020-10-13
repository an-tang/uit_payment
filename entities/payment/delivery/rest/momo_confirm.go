package http

import (
	"net/http"
	"strconv"

	"uit_payment/entities/payment/delivery/request"
	repository "uit_payment/entities/payment/repository"
	service "uit_payment/entities/payment/service"
	paymentRequestRepo "uit_payment/entities/payment_request/repository"
	"uit_payment/enum"
	"uit_payment/lib/env"
	"uit_payment/lib/hmac"
	httpclient "uit_payment/lib/http_client"
	log "uit_payment/lib/logging"
	"uit_payment/lib/providers/momo"
	"uit_payment/model"

	"github.com/sirupsen/logrus"
)

type MomoConfirm struct {
	Handler
	PaymentService     service.PaymentServiceInterface
	PaymentRepo        repository.PaymentRepositoryInterface
	PaymentRequestRepo paymentRequestRepo.PaymentRequestRepositoryInterface
	HTTPClient         httpclient.HTTPCInterface
}

func NewMomoConfirm() *MomoConfirm {
	return &MomoConfirm{
		PaymentRepo:        repository.NewPaymentRepository(),
		PaymentRequestRepo: paymentRequestRepo.NewPaymentRequestRepository(),
		PaymentService:     service.NewPaymentService(),
		HTTPClient:         httpclient.HTTPInstance(),
	}
}

func (m *MomoConfirm) Handle() {
	params := &momo.MomoAIOConfirmRequest{}
	m.ParseParam(&params)

	payment := &model.Payment{}
	err := m.PaymentRepo.FindByPaymentTX(params.RequestID, payment)
	if err != nil {
		log.WithError(err).Errorln("MomoConfirm.CannotFindTransaction: " + params.RequestID)
		return
	}

	req := parseConfirm(*params)
	m.HTTPClient.Post(env.GetMomoNotifyURL(), "application/json", &req, nil)

	paymentRequestLog := &model.PaymentRequest{
		RequestType: enum.PaymentRequestTypeWebhook,
		PaymentID:   payment.ID,
	}

	paymentRequestLog.Populate(params, req, http.StatusOK)
	go m.callbackPartner(params, payment)
	go m.updatePayment(payment, paymentRequestLog, params)
}

func parseConfirm(req momo.MomoAIOConfirmRequest) momo.MomoAIOConfirmResponse {
	resp := momo.MomoAIOConfirmResponse{
		AccessKey:    req.AccessKey,
		PartnerCode:  req.PartnerCode,
		OrderID:      req.OrderID,
		RequestID:    req.RequestID,
		ResponseTime: req.ResponseTime,
		ErrorCode:    0,
		Message:      "Success",
		ExtraData:    "",
	}

	hmacData := resp.HmacCombine()
	resp.Signature = hmac.HexStringEncode(hmac.SHA256, env.GetMomoSecretKey(), hmacData)
	return resp
}

func (m *MomoConfirm) callbackPartner(params *momo.MomoAIOConfirmRequest, payment *model.Payment) {
	amount, err := strconv.ParseFloat(params.Amount, 32)
	if err != nil {
		logrus.WithError(err).Errorln("MomoConfirm.ConvertAmountError:", err.Error())
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

	err = m.HTTPClient.Post(env.UitTravelURL()+"/callback", "application/json", req, nil)
	if err != nil {
		logrus.WithError(err).Errorln("MomoConfirm.CallbackPartnerError:", err.Error())
	}
}

func (m *MomoConfirm) updatePayment(payment *model.Payment, paymentRequestLog *model.PaymentRequest, params *momo.MomoAIOConfirmRequest) {
	if params.ErrorCode != 0 {
		m.PaymentService.UpdateFailed(payment, paymentRequestLog)
	} else {
		m.PaymentService.UpdatePaid(payment, paymentRequestLog)
	}
}
