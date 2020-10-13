package handler

import (
	"context"
	"encoding/json"

	"uit_payment/entities/payment/delivery/request"
	"uit_payment/entities/payment/delivery/response"
	"uit_payment/entities/payment/service"
	"uit_payment/enum"
	"uit_payment/model"
	payment_api "uit_payment/services/payment"

	"github.com/sirupsen/logrus"
)

type PaymentgRPCHandler struct {
	PaymentService service.PaymentServiceInterface
}

func NewPaymentgRPCHandler() payment_api.PaymentServiceServer {
	return &PaymentgRPCHandler{
		PaymentService: service.NewPaymentService(),
	}
}

func (h *PaymentgRPCHandler) CreatePayment(ctx context.Context, req *payment_api.CreatePaymentRequest) (*payment_api.CreatePaymentResponse, error) {
	createPaymentReq := parseParamsCreatePayment(req)
	payment := &model.Payment{}
	a, _ := json.Marshal(createPaymentReq)
	logrus.Warning(string(a))

	payment, err := h.PaymentService.CreatePayment(&createPaymentReq, payment)
	if err != nil {
		return &payment_api.CreatePaymentResponse{}, err
	}

	resp := response.PopulategRPCCreatePayment(*payment)
	return &resp, nil
}

func parseParamsCreatePayment(req *payment_api.CreatePaymentRequest) request.CreatePaymentRequest {
	return request.CreatePaymentRequest{
		TransactionID: req.TransactionId,
		Amount:        req.Amount,
		PartnerKey:    req.PartnerKey,
		PaymentMethod: enum.PaymentMethodMapping(int(req.PaymentMethod)),
		StoreID:       req.StoreId,
		Product:       req.TransactionId,
		Token:         req.Token,
	}
}
