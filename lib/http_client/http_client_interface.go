package httpclient

import "uit_payment/model"

type HTTPCInterface interface {
	PostForm(endpoint string, params interface{}) ([]byte, error)
	GetForm(endpoint string, params interface{}) ([]byte, error)
	SendOrderCallback(payment *model.Payment) error
	Post(endpoint string, contentType string, req interface{}, resp interface{}) error
}
