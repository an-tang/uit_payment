package response

import (
	"uit_payment/enum"
	"uit_payment/model"
	payment_api "uit_payment/services/payment"
)

type Error struct {
	Message string `json:"error"`
	// Code    ErrorCode `json:"code"`
}

type CreatePaymentResponse struct {
	QrText        string             `json:"qr_text"`
	TransactionID string             `json:"transaction_id"`
	PaymentMethod enum.PaymentMethod `json:"payment_method"`
	Status        string             `json:"status"`
	StatusValue   enum.PaymentStatus `json:"status_value"`
}

type HealthResponse struct {
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

func (p *CreatePaymentResponse) PopulateFromModel(obj model.Payment) {
	p.QrText = obj.QrCode
	p.TransactionID = obj.TransactionID
	p.PaymentMethod = obj.PaymentMethod
	p.Status = obj.Status.String()
	p.StatusValue = obj.Status
}

func PopulategRPCCreatePayment(payment model.Payment) payment_api.CreatePaymentResponse {
	return payment_api.CreatePaymentResponse{
		QrText:        payment.QrCode,
		TransactionId: payment.TransactionID,
		PaymentMethod: int32(payment.PaymentMethod),
		Status:        payment.Status.String(),
		StatusValue:   int32(payment.Status),
	}
}
