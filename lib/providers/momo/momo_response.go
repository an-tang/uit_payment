package momo

import "fmt"

type ScanQRCodeResponse struct {
	Status       int     `json:"status"`
	Message      string  `json:"message"`
	PartnerRefID string  `json:"partnerRefId"`
	MomoTransID  string  `json:"momoTransId"`
	Amount       float32 `json:"amount"`
	Signature    string  `json:"signature"`
}

type MomoNotifyURLResponse struct {
	Status       string  `json:"status"`
	Message      string  `json:"success"`
	PartnerRefID string  `json:"partnerRefId"`
	MomoTransID  string  `json:"momoTransId"`
	Amount       float32 `json:"amount"`
	Signature    string  `json:"signature"`
}

type MomoPaymentResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		PartnerCode  string  `json:"partnerCode"`
		PartnerRefID string  `json:"partnerRefId"`
		MomoTransID  string  `json:"momoTransId"`
		Amount       float32 `json:"amount"`
	} `json:"data"`
	Signature string `json:"signature"`
}

type MomoGetPaymentResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Status         int     `json:"status"`
		Message        string  `json:"message"`
		PartnerCode    string  `json:"partnerCode"`
		BillID         string  `json:"billId"`
		TransID        int64   `json:"transId"`
		Amount         float32 `json:"amount"`
		DiscountAmount float32 `json:"discountAmount"`
		Fee            float32 `json:"fee"`
		PhoneNumber    string  `json:"phoneNumber"`
		CustomerName   string  `json:"customerName"`
		StoreID        string  `json:"storeId"`
		RequestDate    string  `json:"requestDate"`
		ResponseDate   string  `json:"responseDate"`
	} `json:"data"`
}

type MomoRefundResponse struct {
	Status       int     `json:"status"`
	Message      string  `json:"message"`
	PartnerRefID string  `json:"partnerRefId"`
	TransID      string  `json:"transId"`
	Amount       float32 `json:"amount"`
}

type MomoCreateAIOResponse struct {
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
