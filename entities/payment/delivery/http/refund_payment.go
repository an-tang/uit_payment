package http

import (
	"uit_payment/entities/payment/delivery/response"
	service "uit_payment/entities/payment/service"
	"uit_payment/model"
)

type RefundPayment struct {
	Handler
	PaymentService service.PaymentServiceInterface
}

func NewRefundPayment() *RefundPayment {
	return &RefundPayment{
		PaymentService: service.NewPaymentService(),
	}
}

func (rp *RefundPayment) Handle() {
	param := rp.Vars()
	transactionID := param["transaction_id"]

	mpayment := &model.Payment{}
	mpayment, err := rp.PaymentService.RefundPayment(transactionID)
	if err != nil {
		rp.RenderError(err.Error())
		return
	}

	resp := &response.CreatePaymentResponse{}
	resp.PopulateFromModel(*mpayment)
	rp.RenderSuccess(resp)
}
