package momo

import "fmt"

type MomoCreatePaymentResponse struct {
	RequestID        string `json:"requestId"`
	ErrorCode        int    `json:"errorCode"`
	OrderID          string `json:"orderId"`
	Message          string `json:"message"`
	LocalMessage     string `json:"localMessage"`
	RequestType      string `json:"requestType"`
	PayURL           string `json:"payUrl"`
	Signature        string `json:"signature"`
	QrCodeURL        string `json:"qrCodeUrl"`
	Deeplink         string `json:"deeplink"`
	DeeplinkWebInApp string `json:"deeplinkWebInApp"`
}
type MomoGetPaymentResponse struct {
	PartnerCode  string `json:"partnerCode"`
	AccessKey    string `json:"accessKey"`
	RequestID    string `json:"requestId"`
	OrderID      string `json:"orderId"`
	RequestType  string `json:"requestType"`
	ExtraData    string `json:"extraData"`
	Amount       string `json:"amount"`
	TransID      string `json:"transId"`
	PayType      string `json:"payType"`
	ErrorCode    int    `json:"errorCode"`
	Message      string `json:"message"`
	LocalMessage string `json:"localMessage"`
	Signature    string `json:"signature"`
}

type MomoRefundResponse struct {
	PartnerCode  string `json:"partnerCode"`
	AccessKey    string `json:"accessKey"`
	RequestID    string `json:"requestId"`
	Amount       string `json:"amount"`
	OrderID      string `json:"orderId"`
	TransID      string `json:"transId"`
	RequestType  string `json:"requestType"`
	Signature    string `json:"signature"`
	ErrorCode    int    `json:"errorCode"`
	Message      string `json:"message"`
	LocalMessage string `json:"localMessage"`
}

type MomoAIOConfirmResponse struct {
	PartnerCode  string `json:"partnerCode"`
	AccessKey    string `json:"accessKey"`
	RequestID    string `json:"requestID"`
	OrderID      string `json:"orderID"`
	ErrorCode    int    `json:"errorCode"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	ExtraData    string `json:"extraData"`
	Signature    string `json:"signature"`
}

func (m MomoAIOConfirmResponse) HmacCombine() string {
	return fmt.Sprintf("partnerCode=%s&accessKey=%s&requestId=%s&orderId=%s&errorCode=%d&message=%s&responseTime=%s&extraData=%s",
		m.PartnerCode, m.AccessKey, m.RequestID, m.OrderID, m.ErrorCode, m.Message, m.ResponseTime, m.ExtraData)
}
