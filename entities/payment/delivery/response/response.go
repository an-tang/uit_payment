package response

import (
	"uit_payment/enum"
	"uit_payment/model"
)

type Error struct {
	Message string `json:"error"`
	// Code    ErrorCode `json:"code"`
}

type Payment struct {
	QrText        string             `json:"qr_text"`
	TransactionID string             `json:"transaction_id"`
	PaymentMethod enum.PaymentMethod `json:"payment_method"`
	Status        string             `json:"status"`
	StatusValue   enum.PaymentStatus `json:"status_value"`
}

func (p *Payment) PopulateFromModel(obj model.Payment) {
	p.QrText = obj.QrCode
	p.TransactionID = obj.TransactionID
	p.PaymentMethod = obj.PaymentMethod
	p.Status = obj.Status.String()
	p.StatusValue = obj.Status
}
