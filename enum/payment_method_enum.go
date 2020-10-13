package enum

type PaymentMethod int

const (
	Momo PaymentMethod = iota + 200
)

func PaymentMethodValue(pm PaymentMethod) string {
	switch pm {

	case Momo:
		return "Momo"
	default:
		return ""
	}
}

func PaymentMethodMapping(e int) PaymentMethod {
	switch e {
	case 200:
		return Momo
	default:
		return Momo
	}
}
