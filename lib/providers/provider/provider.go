package provider

import (
	"uit_payment/enum"
	"uit_payment/lib/providers"
	"uit_payment/lib/providers/momo"
	"uit_payment/model"
)

type provider struct {
}

func NewProvider() providers.Provider {
	provider := &provider{}
	return provider
}

func (p *provider) GetProvider(mpayment *model.Payment) providers.ProviderInterface {
	var Provider = map[enum.PaymentMethod]providers.ProviderInterface{
		enum.Momo: momo.NewClient(),
	}
	return Provider[mpayment.PaymentMethod]
}
