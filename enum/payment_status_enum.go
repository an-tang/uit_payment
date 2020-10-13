package enum

type PaymentStatus int

const (
	PaymentStatusNew PaymentStatus = iota + 1
	PaymentStatusPaid
	PaymentStatusRefund
	PaymentStatusFailed
)

func PaymentStatusValue(str string) PaymentStatus {
	switch str {
	case "NEW":
		return PaymentStatusNew
	case "PAID", "SUCCESS":
		return PaymentStatusPaid
	case "REFUND":
		return PaymentStatusRefund
	case "FAILED":
		return PaymentStatusFailed
	default:
		return 0
	}
}

func (e PaymentStatus) String() string {
	return [...]string{"unknown", "NEW", "PAID", "REFUND", "FAILED", "WAITING_FOR_PAYMENT"}[e]
}
