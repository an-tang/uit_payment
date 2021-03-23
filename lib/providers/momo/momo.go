package momo

import (
	"errors"
	"fmt"
	"net/http"

	"uit_payment/entities/payment/delivery/request"
	"uit_payment/enum"
	"uit_payment/lib/env"
	"uit_payment/lib/hmac"
	httpclient "uit_payment/lib/http_client"
	"uit_payment/lib/providers"
	"uit_payment/model"
)

const (
	REQUEST_TYPE_CREATE_PAYMENT = "captureMoMoWallet"
	REQUEST_TYPE_GET_PAYMENT    = "transactionStatus"
	REQUEST_TYPE_REFUND_PAYMENT = "refundMoMoWallet"
)

type MoMoPayment struct {
	MoMoPartnerCode string
	MoMoAccessKey   string
	MoMoSecretKey   string
	MoMoNotifyURL   string
	MoMoEndpoint    string
	HTTPClient      httpclient.HTTPCInterface
}

func NewClient() providers.ProviderInterface {
	return &MoMoPayment{
		MoMoPartnerCode: env.GetMomoPartnerCode(),
		MoMoAccessKey:   env.GetMomoAccessKey(),
		MoMoSecretKey:   env.GetMomoSecretKey(),
		MoMoNotifyURL:   env.GetMomoCallbackURL(),
		MoMoEndpoint:    env.GetMomoAIOURL(),
		HTTPClient:      httpclient.HTTPInstance(),
	}
}

func (mp *MoMoPayment) Name() string {
	return enum.PaymentMethodValue(enum.Momo)
}

func (mp *MoMoPayment) CreatePayment(paymentRequest *request.CreatePaymentRequest, paymentModel *model.Payment) (*model.PaymentRequest, error) {
	paymentRequestLog := &model.PaymentRequest{RequestType: enum.PaymentRequestTypeCreate}
	req := parseParamToCreateAIOPayment(paymentRequest, paymentModel)
	resp := &MomoCreatePaymentResponse{}

	err := mp.HTTPClient.Post(mp.MoMoEndpoint, "application/json", req, resp)
	if err != nil {
		paymentRequestLog.Populate(req, resp, http.StatusBadRequest)
		return paymentRequestLog, err
	}

	paymentModel.QrCode = resp.PayURL
	paymentModel.PaymentTX = resp.RequestID
	paymentRequestLog.Populate(req, resp, http.StatusOK)
	return paymentRequestLog, nil
}

func (mp *MoMoPayment) GetPayment(paymentModel *model.Payment) (*model.PaymentRequest, error) {
	paymentRequestLog := &model.PaymentRequest{
		RequestType: enum.PaymentRequestTypeGetDetail,
		PaymentID:   paymentModel.ID,
	}

	req := mp.parseGetPaymentRequest(paymentModel)
	resp := &MomoGetPaymentResponse{}

	err := mp.HTTPClient.Post(mp.MoMoEndpoint, "Application/json", req, resp)
	if err != nil {
		paymentRequestLog.Populate(req, resp, http.StatusBadRequest)
		return paymentRequestLog, err
	}

	if resp.ErrorCode != 0 {
		paymentRequestLog.Populate(req, resp, http.StatusBadRequest)
		return paymentRequestLog, errors.New(resp.Message)
	}

	paymentModel.PaymentTX = resp.TransID
	paymentRequestLog.Populate(req, resp, http.StatusOK)

	return paymentRequestLog, nil
}

func (mp *MoMoPayment) RefundPayment(paymentModel *model.Payment) (*model.PaymentRequest, error) {
	paymentRequestLog := &model.PaymentRequest{
		RequestType: enum.PaymentRequestTypeRefund,
	}

	req := mp.parseRefundPaymentRequest(paymentModel)
	resp := &MomoRefundResponse{}

	err := mp.HTTPClient.Post(mp.MoMoEndpoint, "Application/json", req, resp)
	if err != nil {
		paymentRequestLog.Populate(req, resp, http.StatusBadRequest)
		return paymentRequestLog, err
	}

	if resp.ErrorCode != 0 {
		paymentRequestLog.Populate(req, resp, http.StatusBadRequest)
		return paymentRequestLog, errors.New(resp.Message)
	}

	paymentRequestLog.Populate(req, resp, http.StatusOK)
	return paymentRequestLog, nil
}

func parseParamToCreateAIOPayment(paymentRequest *request.CreatePaymentRequest, paymentModel *model.Payment) MomoCreatePaymentRequest {
	request := MomoCreatePaymentRequest{
		AccessKey:   env.GetMomoAccessKey(),
		PartnerCode: env.GetMomoPartnerCode(),
		RequestID:   paymentModel.GenerateUID(),
		Amount:      fmt.Sprintf("%v", int(paymentModel.Amount)),
		NotifyURL:   fmt.Sprintf("%s%s", env.GetMomoCallbackURL(), "/momo/confirm"),
		OrderID:     paymentModel.TransactionID,
		OrderInfo:   "Thanh toán đơn hàng UIT Shop",
		RequestType: "captureMoMoWallet",
		ReturnURL:   paymentRequest.RedirectURL,
	}

	hmacData := request.CombineHmacData()
	request.Signature = hmac.HexStringEncode(hmac.SHA256, env.GetMomoSecretKey(), hmacData)

	return request
}

func (mp *MoMoPayment) parseGetPaymentRequest(paymentModel *model.Payment) MomoGetPaymentRequest {
	request := MomoGetPaymentRequest{
		AccessKey:   mp.MoMoAccessKey,
		PartnerCode: mp.MoMoPartnerCode,
		RequestID:   paymentModel.GenerateUID(),
		OrderID:     paymentModel.TransactionID,
		RequestType: "transactionStatus",
	}

	hmacData := request.CombineHmacData()
	request.Signature = hmac.HexStringEncode(hmac.SHA256, mp.MoMoSecretKey, hmacData)

	return request
}

func (mp *MoMoPayment) parseRefundPaymentRequest(paymentModel *model.Payment) MomoRefundPaymentRequest {
	request := MomoRefundPaymentRequest{
		AccessKey:   mp.MoMoAccessKey,
		PartnerCode: mp.MoMoPartnerCode,
		RequestID:   paymentModel.GenerateUID(),
		Amount:      fmt.Sprintf("%d", int(paymentModel.Amount)),
		OrderID:     fmt.Sprintf("REFUND_%s", paymentModel.TransactionID),
		TransID:     paymentModel.PaymentTX,
		RequestType: REQUEST_TYPE_REFUND_PAYMENT,
	}

	hmacData := request.CombineHmacData()
	request.Signature = hmac.HexStringEncode(hmac.SHA256, mp.MoMoSecretKey, hmacData)

	return request
}
