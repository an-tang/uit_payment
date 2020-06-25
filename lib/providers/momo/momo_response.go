package momo

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
