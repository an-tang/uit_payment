package model

import (
	"fmt"
	"time"

	"uit_payment/enum"
)

type Payment struct {
	BaseModel
	UID           string             `gorm:"column:uid" json:"uid"`
	TransactionID string             `gorm:"column:transaction_id" json:"transaction_id"`
	PaymentMethod enum.PaymentMethod `gorm:"column:payment_method" json:"payment_method"`
	Amount        float32            `gorm:"column:amount" json:"amount"`
	Currency      string             `gorm:"column:currency" json:"currency"`
	StoreID       string             `gorm:"column:store_id" json:"store_id"`
	QrCode        string             `gorm:"column:qr_code" json:"qr_code"`
	PaymentTX     string             `gorm:"column:payment_tx" json:"payment_tx"`
	Status        enum.PaymentStatus `gorm:"column:status" json:"status"`
	PaidAt        *time.Time         `gorm:"column:paid_at" json:"paid_at"`
	CompletedAt   *time.Time         `gorm:"column:completed_at" json:"completed_at"`
}

// Set table name
func (Payment) TableName() string {
	return "payments"
}

func (payment Payment) GenerateUID() string {
	now := time.Now()
	yyMMdd := fmt.Sprintf("%02d%02d%02d", now.Year()%100, int(now.Month()), now.Day())
	return fmt.Sprintf("%v_%v", yyMMdd, payment.TransactionID)
}
