package dom_api

import (
	"context"

	"uit_payment/model"
)

type DOMServiceInterface interface {
	PaymentCallback(c context.Context, payment model.Payment) error
}
