package enum

type PaymentMethod int

const (
	Momo PaymentMethod = iota + 100
)

func PaymentMethodValue(pm PaymentMethod) string {
	switch pm {

	case Momo:
		return "Momo"
	default:
		return ""
	}
}
