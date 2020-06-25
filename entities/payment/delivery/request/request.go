package request

import "uit_payment/enum"

type CreatePaymentRequest struct {
	TransactionID string             `json:"transaction_id"`
	Amount        float32            `json:"amount"`
	Currency      string             `json:"currency"`
	PaymentMethod enum.PaymentMethod `json:"payment_method"`
	StoreID       string             `json:"store_id"`
}
