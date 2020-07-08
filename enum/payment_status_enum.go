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
	case "new":
		return PaymentStatusNew
	case "paid", "success":
		return PaymentStatusPaid
	case "refund":
		return PaymentStatusRefund
	case "failed", "pending", "bad_debt":
		return PaymentStatusFailed
	default:
		return 0
	}
}

func (e PaymentStatus) String() string {
	return [...]string{"unknown", "new", "paid", "refund", "failed", "waiting_for_payment"}[e]
}
