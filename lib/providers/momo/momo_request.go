package momo

import "fmt"

type MomoOrderRequest struct {
	Amount        float32 `json:"amount"`
	StoreID       string  `json:"storeId"`
	TransactionID string  `json:"transaction_id"`
	Signature     string  `json:"signature"`
	Domain        string  `json:"domain"`
	PartnerCode   string  `json:"partnerCode"`
	StoreSlug     string  `json:"storeSlug"`
}

type MomoFieldNotifyURLRequest struct {
	PartnerCode    string  `json:"partnerCode"`
	AccessKey      string  `json:"accessKey"`
	Amount         float32 `json:"amount"`
	PartnerRefID   string  `json:"partnerRefId"`
	PartnerTransID string  `json:"partnerTransId"`
	TransType      string  `json:"transType"`
	MomoTransID    string  `json:"momoTransId"`
	Status         int     `json:"status"`
	Message        string  `json:"message"`
	ResponseTime   int64   `json:"responseTime"`
	StoreID        string  `json:"storeId"`
	Signature      string  `json:"signature"`
}

type MomoCreateAIORequest struct {
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

type MomoAIOConfirmRequest struct {
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

type MomoPaymentRequest struct {
	PartnerCode    string `json:"partnerCode"`
	PartnerRefID   string `json:"partnerRefId"`
	RequestType    string `json:"requestType"`
	RequestID      string `json:"requestId"`
	MomoTransID    string `json:"momoTransId"`
	Signature      string `json:"signature"`
	CustomerNumber string `json:"customerNumber"`
	Description    string `json:"description"`
}

type MomoGetPaymentRequest struct {
	PartnerCode  string  `json:"partnerCode"`
	PartnerRefID string  `json:"partnerRefId"`
	Hash         string  `json:"hash"`
	Version      float32 `json:"version"`
	MomoTransID  string  `json:"momoTransId"`
}

type MomoRefundRequest struct {
	PartnerCode string  `json:"partnerCode"`
	RequestID   string  `json:"requestId"`
	Hash        string  `json:"hash"`
	Version     float32 `json:"version"`
}

type HashGetPayment struct {
	PartnerCode  string `json:"partnerCode"`
	PartnerRefID string `json:"partnerRefId"`
	RequestID    string `json:"requestId"`
	MomoTransID  string `json:"momoTranId"`
}

type HashRefundPayment struct {
	PartnerCode  string      `json:"partnerCode"`
	PartnerRefID string      `json:"partnerRefId"`
	MomoTransID  string      `json:"momoTransId"`
	Amount       float32     `json:"amount"`
	StoreID      string      `json:"storeId"`
	Description  string      `json:"description"`
	Extra        interface{} `json:"extra"`
}

func (m *MomoOrderRequest) HmacCombine() string {
	return fmt.Sprintf("storeSlug=%s&amount=%v&billId=%s", m.StoreSlug, int64(m.Amount), m.TransactionID)
}

func (m *MomoNotifyURLResponse) HmacCombine() string {
	return fmt.Sprintf("amount=%v&message=%s&momoTransId=%s&partnerRefId=%v&status=%v",
		m.Amount, m.Message, m.MomoTransID, m.PartnerRefID, m.Status)
}

func (m *MomoPaymentRequest) HmacCombine() string {
	return fmt.Sprintf("partnerCode=%s&partnerRefId=%s&requestType=%s&requestId=%s&momoTransId=%s",
		m.PartnerCode, m.PartnerRefID, m.RequestType, m.RequestID, m.MomoTransID)
}

func (m *MomoCreateAIORequest) HmacCombine() string {
	return fmt.Sprintf("partnerCode=%s&accessKey=%s&requestId=%s&amount=%s&orderId=%s&orderInfo=%s&returnUrl=%s&notifyUrl=%s&extraData=%s",
		m.PartnerCode, m.AccessKey, m.RequestID, m.Amount, m.OrderID, m.OrderInfo, m.ReturnURL, m.NotifyURL, m.ExtraData)
}

func (m *MomoAIOConfirmRequest) HmacCombine() string {
	return fmt.Sprintf("partnerCode=%s&accessKey=%s&requestId=%s&amount=%s&orderId=%s&orderInfo=%s&orderType=%s&transId=%s&message=%s&localMessage=%s&responseTime=%s&errorCode=%d&payType=%s&extraData=%s",
		m.PartnerCode, m.AccessKey, m.RequestID, m.Amount, m.OrderID, m.OrderInfo, m.OrderType, m.TransID, m.Message, m.LocalMessage, m.ResponseTime, m.ErrorCode, m.PayType, m.ExtraData)
}
