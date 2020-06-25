package http

import (
	"uit_payment/entities/payment/delivery/response"
	service "uit_payment/entities/payment/service"
	"uit_payment/model"
)

type GetPayment struct {
	Handler
	PaymentService service.PaymentServiceInterface
}

func NewGetPayment() *GetPayment {
	return &GetPayment{
		PaymentService: service.NewPaymentUsecase(),
	}
}

func (gp *GetPayment) Handle() {
	param := gp.Vars()
	transactionID := param["transaction_id"]

	payment := &model.Payment{}
	payment, err := gp.PaymentService.GetPayment(transactionID)
	if err != nil {
		gp.RenderError(err.Error())
		return
	}

	resp := &response.Payment{}
	resp.PopulateFromModel(*payment)
	gp.RenderSuccess(resp)
}
