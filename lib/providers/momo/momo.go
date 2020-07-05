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
	log "uit_payment/lib/logging"
	"uit_payment/lib/providers"
	"uit_payment/lib/rsa"
	"uit_payment/model"
)

type MomoPayment struct {
	MomoQRCodeURL   string
	MomoPartnerCode string
	MomoPublicKey   string
	HTTPClient      httpclient.HTTPCInterface
}

func NewClient() providers.ProviderInterface {
	return &MomoPayment{
		MomoQRCodeURL:   env.GetMomoQRCodeURL(),
		MomoPartnerCode: env.GetMomoPartnerCode(),
		MomoPublicKey:   env.GetMomoPublicKey(),
		HTTPClient:      httpclient.HTTPInstance(),
	}
}

func (mp *MomoPayment) Name() string {
	return enum.PaymentMethodValue(enum.Momo)
}

// func (mp *MomoPayment) CreatePayment(paymentRequest *request.CreatePaymentRequest, paymentModel *model.Payment) (*model.PaymentRequest, error) {
// 	paymentRequestLog := &model.PaymentRequest{RequestType: enum.PaymentRequestGenerateQRCode}
// 	momoReq := mp.parseParamToCreatePaymentRequest(paymentRequest, paymentModel)
// 	qrCode, err := generateQRCode(momoReq)
// 	if err != nil {
// 		paymentRequestLog.Populate(momoReq, nil, http.StatusBadRequest)
// 		return paymentRequestLog, err
// 	}

// 	paymentModel.QrCode = qrCode
// 	paymentModel.UID = paymentModel.GenerateUID()
// 	paymentRequestLog.Populate(momoReq, nil, http.StatusOK)
// 	return paymentRequestLog, nil
// }

func (mp *MomoPayment) CreatePayment(paymentRequest *request.CreatePaymentRequest, paymentModel *model.Payment) (*model.PaymentRequest, error) {
	paymentRequestLog := &model.PaymentRequest{RequestType: enum.PaymentRequestTypeCreate}
	momoReq := parseParamToCreateAIOPayment(paymentRequest, paymentModel)
	momoResp := &MomoCreateAIOResponse{}
	endpoint := env.GetMomoCreateAIOURL()

	err := mp.HTTPClient.Post(endpoint, "application/json", momoReq, momoResp)
	if err != nil {
		paymentRequestLog.Populate(momoReq, momoResp, http.StatusBadRequest)
		return paymentRequestLog, err
	}

	paymentModel.QrCode = momoResp.PayURL
	paymentModel.PaymentTX = momoResp.RequestID
	return paymentRequestLog, nil
}

func (mp *MomoPayment) GetPayment(paymentModel *model.Payment) (*model.PaymentRequest, error) {
	paymentRequestLog := &model.PaymentRequest{
		RequestType: enum.PaymentRequestTypeGetDetail,
		PaymentID:   paymentModel.ID,
	}

	getPaymentReq, err := mp.parseParamToGetRequest(paymentModel)
	if err != nil {
		return paymentRequestLog, err
	}

	endpoint := env.GetMomoGetPaymentURL()
	getPaymentResp := &MomoGetPaymentResponse{}
	err = mp.HTTPClient.Post(endpoint, "application/json", getPaymentReq, getPaymentResp)
	if err != nil {
		paymentRequestLog.Populate(getPaymentReq, getPaymentResp, http.StatusBadRequest)
		return paymentRequestLog, err
	}

	paymentRequestLog.Populate(getPaymentReq, getPaymentResp, http.StatusOK)
	return paymentRequestLog, nil
}

func (mp *MomoPayment) RefundPayment(paymentModel *model.Payment) (*model.PaymentRequest, error) {
	paymentRequestLog := &model.PaymentRequest{
		RequestType: enum.PaymentRequestTypeRefund,
	}

	refundPaymentReq, err := mp.parseParamToRefundRequest(paymentModel)
	if err != nil {
		return paymentRequestLog, err
	}

	endpoint := env.GetMomoRefundURL()
	refundPaymentResp := &MomoRefundResponse{}
	err = mp.HTTPClient.Post(endpoint, "application/json", refundPaymentReq, refundPaymentResp)
	if err != nil {
		paymentRequestLog.Populate(refundPaymentReq, refundPaymentResp, http.StatusBadRequest)
		return paymentRequestLog, err
	}

	if refundPaymentResp.Status != 0 {
		return paymentRequestLog, errors.New(refundPaymentResp.Message)
	}

	paymentModel.Status = enum.PaymentStatusRefund
	paymentRequestLog.Populate(refundPaymentReq, refundPaymentResp, http.StatusOK)
	return paymentRequestLog, nil
}

func generateQRCode(order *MomoOrderRequest) (string, error) {
	if order.Domain == "" {
		log.WithError(errors.New("Cannot find momo domain")).Errorln("Momo generate qr code error: " + order.TransactionID)
		return "", errors.New("Cannot find momo domain")
	}

	hmacData := order.HmacCombine()
	order.Signature = hmac.HexStringEncode(hmac.SHA256, env.GetMomoSecretKey(), hmacData)
	return fmt.Sprintf("%s/%s?a=%v&b=%s&s=%s", order.Domain, order.StoreSlug, int64(order.Amount), order.TransactionID, order.Signature), nil
}

func (mp *MomoPayment) parseParamToCreatePaymentRequest(paymentRequest *request.CreatePaymentRequest, paymentModel *model.Payment) *MomoOrderRequest {
	momoReq := &MomoOrderRequest{
		Amount:        paymentRequest.Amount,
		TransactionID: paymentRequest.TransactionID,
		StoreID:       paymentRequest.StoreID,
		Domain:        mp.MomoQRCodeURL,
		PartnerCode:   mp.MomoPartnerCode,
	}

	momoReq.StoreSlug = fmt.Sprintf("%s-%s", momoReq.PartnerCode, momoReq.StoreID)
	paymentModel.Amount = paymentRequest.Amount
	paymentModel.StoreID = paymentRequest.StoreID
	return momoReq
}

func (mp *MomoPayment) parseParamToGetRequest(paymentModel *model.Payment) (MomoGetPaymentRequest, error) {
	hashParam := HashGetPayment{
		PartnerCode:  mp.MomoPartnerCode,
		PartnerRefID: paymentModel.TransactionID,
		RequestID:    paymentModel.UID,
	}

	hash, err := rsa.CreateRSA(hashParam, mp.MomoPublicKey)
	if err != nil {
		return MomoGetPaymentRequest{}, err
	}

	getPaymentReq := &MomoGetPaymentRequest{
		PartnerCode:  mp.MomoPartnerCode,
		PartnerRefID: paymentModel.TransactionID,
		Hash:         hash,
		Version:      env.GetMomoVersion(),
	}
	return *getPaymentReq, nil
}

func (mp *MomoPayment) parseParamToRefundRequest(paymentModel *model.Payment) (MomoRefundRequest, error) {
	hashParam := HashRefundPayment{
		PartnerCode:  mp.MomoPartnerCode,
		PartnerRefID: paymentModel.TransactionID,
		MomoTransID:  paymentModel.PaymentTX,
		Amount:       paymentModel.Amount,
	}

	hash, err := rsa.CreateRSA(hashParam, mp.MomoPublicKey)
	if err != nil {
		return MomoRefundRequest{}, err
	}

	refundPaymentReq := &MomoRefundRequest{
		PartnerCode: mp.MomoPartnerCode,
		RequestID:   paymentModel.UID,
		Hash:        hash,
		Version:     env.GetMomoVersion(),
	}
	return *refundPaymentReq, nil
}

func parseParamToCreateAIOPayment(paymentRequest *request.CreatePaymentRequest, paymentModel *model.Payment) MomoCreateAIORequest {
	request := MomoCreateAIORequest{
		AccessKey:   env.GetMomoAccessKey(),
		PartnerCode: env.GetMomoPartnerCode(),
		RequestID:   paymentModel.GenerateUID(),
		Amount:      fmt.Sprintf("%v", paymentModel.Amount),
		NotifyURL:   fmt.Sprintf("%s%s", env.MomoCallbackURL(), "/momo/confirm"),
		OrderID:     paymentModel.TransactionID,
		OrderInfo:   paymentRequest.TourName,
		RequestType: "captureMoMoWallet",
		ReturnURL:   env.UitTravelURL(),
	}

	hmacData := request.HmacCombine()
	request.Signature = hmac.HexStringEncode(hmac.SHA256, env.GetMomoSecretKey(), hmacData)

	return request
}
