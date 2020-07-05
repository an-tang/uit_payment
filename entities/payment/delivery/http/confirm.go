package http

import (
	"net/http"
	"strconv"

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

const capture string = "capture"
const revertAuth string = "revertAuthorize"

type Confirm struct {
	Handler
	PaymentService     service.PaymentServiceInterface
	PaymentRepo        repository.PaymentRepositoryInterface
	PaymentRequestRepo paymentRequestRepo.PaymentRequestRepositoryInterface
	HTTPClient         httpclient.HTTPCInterface
}

func NewConfirm() *Confirm {
	return &Confirm{
		PaymentRepo:        repository.NewPaymentRepository(),
		PaymentRequestRepo: paymentRequestRepo.NewPaymentRequestRepository(),
		PaymentService:     service.NewPaymentService(),
		HTTPClient:         httpclient.HTTPInstance(),
	}
}

func (this *Confirm) Handle() {
	logrus.Warning("====Start=====")
	params := &momo.MomoFieldNotifyURLRequest{}
	this.ParseParam(&params)

	payment := &model.Payment{}
	err := this.PaymentRepo.FindByTransactionID(params.PartnerRefID, payment)
	if err != nil {
		log.WithError(err).Errorln("Momo find transaction error: " + params.PartnerRefID)
		return
	}
	this.responseServerMomo(params)

	paymentRequestLog := &model.PaymentRequest{
		RequestType: enum.PaymentRequestTypeCreate,
		PaymentID:   payment.ID,
	}
	paymentRequestLog.Populate(nil, params, http.StatusOK)
	this.PaymentRequestRepo.Create(paymentRequestLog)

	momoPaymentRequest := createPaymentRequest(params, payment)
	endpoint := env.GetMomoConfirmURL()
	momoPaymentResponse := &momo.MomoPaymentResponse{}
	err = this.HTTPClient.Post(endpoint, "application/json", momoPaymentRequest, momoPaymentResponse)
	if err != nil {
		log.WithError(err).Errorln("Momo confirm error: " + momoPaymentRequest.PartnerRefID)
		return
	}

	payment.PaymentTX = momoPaymentResponse.Data.MomoTransID
	paymentRequest := &model.PaymentRequest{
		RequestType: enum.PaymentRequestTypeWebhook,
	}
	paymentRequest.Populate(momoPaymentRequest, momoPaymentResponse, http.StatusOK)
	go this.PaymentService.UpdatePaid(payment, paymentRequest)
}

func (this *Confirm) responseServerMomo(params *momo.MomoFieldNotifyURLRequest) {
	respNotifyURL := &momo.MomoNotifyURLResponse{
		Status:       strconv.FormatInt(int64(params.Status), 10),
		Message:      params.Message,
		Amount:       params.Amount,
		PartnerRefID: params.PartnerRefID,
		MomoTransID:  params.MomoTransID,
	}

	hmacData := respNotifyURL.HmacCombine()
	respNotifyURL.Signature = hmac.HexStringEncode(hmac.SHA256, env.GetMomoSecretKey(), hmacData)
	this.RenderSuccess(respNotifyURL)
}

func validParam(params *momo.MomoFieldNotifyURLRequest, payment *model.Payment) bool {
	if params.PartnerCode == env.GetMomoPartnerCode() && params.Amount == payment.Amount && params.StoreID == payment.StoreID {
		return true
	}
	return false
}

func createPaymentRequest(params *momo.MomoFieldNotifyURLRequest, payment *model.Payment) momo.MomoPaymentRequest {
	momoPaymentReq := momo.MomoPaymentRequest{
		PartnerCode:  params.PartnerCode,
		PartnerRefID: params.PartnerRefID,
		RequestID:    payment.UID,
		MomoTransID:  params.MomoTransID,
	}

	if validParam(params, payment) {
		momoPaymentReq.RequestType = capture // Xác nhận giao dịch
	} else {
		momoPaymentReq.RequestType = revertAuth // Hủy bỏ giao dịch
	}

	hmacData := momoPaymentReq.HmacCombine()
	momoPaymentReq.Signature = hmac.HexStringEncode(hmac.SHA256, env.GetMomoSecretKey(), hmacData)
	return momoPaymentReq
}
