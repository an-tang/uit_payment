package model

import (
	"uit_payment/enum"
)

type ProductsInfo struct {
	ProductID    int32   `json:"product_id"`
	ProductUPC   string  `json:"product_upc"`
	Quantity     int32   `json:"quantity"`
	Amount       float32 `json:"amount"`
	ProductName  string  `json:"product_name"`
	ProductPrice float32 `json:"product_price"`
}

type PaymentRequest struct {
	BaseModel
	PaymentID   int                     `gorm:"column:payment_id" json:"payment_id"`
	Req         JSONB                   `gorm:"column:req" json:"req"`
	Resp        JSONB                   `gorm:"column:resp" json:"resp"`
	Status      int                     `gorm:"column:status" json:"status"`
	RequestType enum.PaymentRequestType `gorm:"column:request_type" json:"request_type"`
}

// Set table name
func (PaymentRequest) TableName() string {
	return "payment_requests"
}

func (pr *PaymentRequest) Populate(Req interface{}, Resp interface{}, status int) {
	pr.Req.Set(Req)
	pr.Resp.Set(Resp)
	pr.Status = status
}
