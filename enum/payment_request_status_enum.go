package enum

type PaymentRequestType int

const (
	PaymentRequestTypeCreate PaymentRequestType = iota + 1
	PaymentRequestTypeWebhook
	PaymentRequestTypeGetDetail
	PaymentRequestTypeRefund
	PaymentRequestGenerateQRCode
)

func PaymentRequestTypeValue(str string) PaymentRequestType {
	switch str {
	case "create":
		return PaymentRequestTypeCreate
	case "webhook":
		return PaymentRequestTypeWebhook
	case "get_detail":
		return PaymentRequestTypeGetDetail
	case "refund":
		return PaymentRequestTypeRefund
	case "generateQRCode":
		return PaymentRequestGenerateQRCode
	default:
		return 0
	}
}

func (e PaymentRequestType) String() string {
	return [...]string{"unknown", "create", "webhook", "get_detail", "refund"}[e]
}
