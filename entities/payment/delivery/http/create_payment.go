package http

import (
	"uit_payment/entities/payment/delivery/request"
	"uit_payment/entities/payment/delivery/response"
	usecase "uit_payment/entities/payment/service"
	"uit_payment/model"
)

type CreatePayment struct {
	Handler
	PaymentService usecase.PaymentServiceInterface
}

func NewCreatePayment() *CreatePayment {
	return &CreatePayment{
		PaymentService: usecase.NewPaymentUsecase(),
	}
}

func (cp *CreatePayment) Handle() {
	param := &request.CreatePaymentRequest{}
	cp.ParseParam(param)

	mpayment := &model.Payment{}
	mpayment, err := cp.PaymentService.CreatePayment(param, mpayment)

	if err != nil {
		cp.RenderError(err.Error())
		return
	}

	resp := &response.Payment{}
	resp.PopulateFromModel(*mpayment)
	cp.RenderSuccess(resp)
}
