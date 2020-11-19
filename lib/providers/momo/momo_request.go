package momo

import "fmt"

type MomoCreatePaymentRequest struct {
	AccessKey   string `json:"accessKey"`
	PartnerCode string `json:"partnerCode"`
	RequestType string `json:"requestType"`
	NotifyURL   string `json:"notifyUrl"`
	ReturnURL   string `json:"returnUrl"`
	OrderID     string `json:"orderId"`
	Amount      string `json:"amount"`
	OrderInfo   string `json:"orderInfo"`
	RequestID   string `json:"requestId"`
	ExtraData   string `json:"extraData"`
	Signature   string `json:"signature"`
}

type MomoConfirmPaymentRequest struct {
	PartnerCode  string `json:"partnerCode" form:"partnerCode"`
	AccessKey    string `json:"accessKey" form:"accessKey"`
	RequestID    string `json:"requestId" form:"requestId"`
	Amount       string `json:"amount" form:"amount"`
	OrderID      string `json:"orderId" form:"orderID"`
	OrderInfo    string `json:"orderInfo" form:"orderInfo"`
	OrderType    string `json:"orderType" form:"orderType"`
	TransID      string `json:"transID" form:"transID"`
	ErrorCode    int    `json:"errorCode" form:"errorCode"`
	Message      string `json:"message" form:"message"`
	LocalMessage string `json:"localMessage" form:"localMessage"`
	PayType      string `json:"payType" form:"payType"`
	ResponseTime string `json:"responseTime" form:"responseTime"`
	ExtraData    string `json:"extraData" form:"extraData"`
	Signature    string `json:"signature" form:"signature"`
}
type MomoGetPaymentRequest struct {
	PartnerCode string `json:"partnerCode"`
	AccessKey   string `json:"accessKey"`
	RequestID   string `json:"requestId"`
	OrderID     string `json:"orderId"`
	RequestType string `json:"requestType"`
	Signature   string `json:"signature"`
}

type MomoRefundPaymentRequest struct {
	PartnerCode string `json:"partnerCode"`
	AccessKey   string `json:"accessKey"`
	RequestID   string `json:"requestId"`
	Amount      string `json:"amount"`
	OrderID     string `json:"orderId"`
	TransID     string `json:"transId"`
	RequestType string `json:"requestType"`
	Signature   string `json:"signature"`
}

func (m *MomoCreatePaymentRequest) CombineHmacData() string {
	return fmt.Sprintf("partnerCode=%s&accessKey=%s&requestId=%s&amount=%s&orderId=%s&orderInfo=%s&returnUrl=%s&notifyUrl=%s&extraData=%s",
		m.PartnerCode, m.AccessKey, m.RequestID, m.Amount, m.OrderID, m.OrderInfo, m.ReturnURL, m.NotifyURL, m.ExtraData)
}

func (m *MomoGetPaymentRequest) CombineHmacData() string {
	return fmt.Sprintf("partnerCode=%s&accessKey=%s&requestId=%s&orderId=%s&requestType=%s",
		m.PartnerCode, m.AccessKey, m.RequestID, m.OrderID, m.RequestType)
}

func (m *MomoRefundPaymentRequest) CombineHmacData() string {
	return fmt.Sprintf("partnerCode=%s&accessKey=%s&requestId=%s&amount=%s&orderId=%s&transId=%s&requestType=%s",
		m.PartnerCode, m.AccessKey, m.RequestID, m.Amount, m.OrderID, m.TransID, m.RequestType)
}
