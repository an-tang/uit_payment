package model

import (
	"uit_payment/enum"
)

type PaymentRequest struct {
	BaseModel
	PaymentID   int                     `gorm:"column:payment_id" json:"payment_id"`
	Req         JSONB                   `gorm:"column:req" json:"req"`
	Resp        JSONB                   `gorm:"column:resp" json:"resp"`
	Status      int                     `gorm:"column:status" json:"status"`
	RequestType enum.PaymentRequestType `gorm:"column:request_type" json:"request_type"`
}

// TableName mapping table name in database
func (PaymentRequest) TableName() string {
	return "payment_requests"
}

func (pr *PaymentRequest) Populate(Req interface{}, Resp interface{}, status int) {
	pr.Req.Set(Req)
	pr.Resp.Set(Resp)
	pr.Status = status
}
