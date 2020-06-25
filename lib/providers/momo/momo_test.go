package momo_test

// import (
// 	"errors"
// 	"net/http"
// 	"os"
// 	"testing"

// 	"payment/enum"
// 	"payment/lib/env"
// 	httpclient "payment/lib/http_client"
// 	"payment/lib/providers/momo"
// 	"payment/mock"
// 	"payment/model"

// 	"github.com/stretchr/testify/assert"
// )

// var publicKey = "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAlHdcz3tSAQH0YWARfVq2TVAEbxvWlqwSAairWFoffEe8ksw2I8d1K01iOnlibdn6tL7z9j528qRxPkHdZppTW08ZQN/uEUqFC488fQkgsE7CkOdaEsiol7w3nQpW/GsZjHjfLec20HrwDHuNz2U+WRIrF4+AS0TPao3aMSljjcF2oFtL9yjBhddrHegBZfMwaaFMfzeD3bl4DGzHkIw2RE0o807F7Eq9PjfXSrqD6yc8l3rp10/QLE6pXpuuUC51uzswSC9PDQM6UbX/9OO95Iprhj2hH3OXCQrmdeqaPwO2hEZdVeUsb+VQSx63Yj9C3vj3WzW348rFW1LbebpXPOOaosAechkNKdi8LA71eG9cqFdQgDH0SSxyg8NURQZPyFyw3+3RI2DYI8y5Ab2qCLAUt7RU055VT0DQkyuCZI5iu04f4qWzp+3HJIduNYF+PFmuznhrfUt+Y1p9I827U6jXKci81zo/yA3fDcf6JLAMF8HvZ0iADvnINkV+TdPzmHliw3ydy7oTk9lP+d+5fQJKv/AMNQ0dK/IsbiGnXcNPobHsQQiMexWcvYGhgp1LVFQghA/KKwzHKNsYuxnIrza3JmqcAea5RGZa95aP2SGRGGTPsEyPOGFCfcgf7jJWw+vrovI1W/pYnQtX992Wq/3zmr27u8CtSPO0tiE5Y6MCAwEAAQ=="

// var mockPaymentRequest = model.CreatePaymentRequest{
// 	TransactionID: "3023921123831000474253",
// 	TerminalID:    "1111",
// 	Amount:        11111,
// 	Currency:      "VND",
// 	GrabID:        "1111",
// 	StoreID:       "1011",
// }

// func setupClient(http httpclient.HTTPInterface) *momo.MomoPayment {
// 	return &momo.MomoPayment{
// 		HTTPClient:      http,
// 		MomoQRCodeURL:   env.GetMomoQRCodeURL(),
// 		MomoPartnerCode: env.GetMomoPartnerCode(),
// 		MomoPublicKey:   env.GetMomoPublicKey(),
// 	}
// }

// func TestCreatePayment(t *testing.T) {
// 	cases := []struct {
// 		Context            string
// 		It                 string
// 		PaymentRequest     model.CreatePaymentRequest
// 		PaymentModel       model.Payment
// 		HTTPClient         httpclient.HTTPInterface
// 		ExpectedPaymentReq model.PaymentRequest
// 		ExpectedError      error
// 		ExpectedQRCode     string
// 	}{
// 		{
// 			Context:            "Momo: test create payment: generate QR code fail",
// 			HTTPClient:         mock.NewHTTPFailed(),
// 			PaymentRequest:     mockPaymentRequest,
// 			PaymentModel:       model.Payment{TransactionID: "00000", Amount: 10000, StoreID: "1011"},
// 			ExpectedPaymentReq: model.PaymentRequest{Status: http.StatusBadRequest},
// 			ExpectedError:      errors.New("Cannot find momo domain"),
// 		},
// 		{
// 			Context:            "Momo: test create payment: success",
// 			HTTPClient:         mock.NewHTTPSuccess(),
// 			PaymentRequest:     mockPaymentRequest,
// 			PaymentModel:       model.Payment{TransactionID: "22222", Amount: 10000},
// 			ExpectedPaymentReq: model.PaymentRequest{Status: http.StatusOK},
// 			ExpectedError:      nil,
// 			ExpectedQRCode:     "https://momo.test.api/ssv_momo-1011?a=11111&b=3023921123831000474253&s=f85d90ebf3b0a0c0c28a01591312659f25c50b5161cfce31bfa0acc51021018f",
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.Context, func(t *testing.T) {
// 			if c.PaymentModel.TransactionID == "00000" {
// 				os.Setenv("MOMO_QRCODE_URL", "")
// 			} else {
// 				os.Setenv("MOMO_QRCODE_URL", "https://momo.test.api")
// 				os.Setenv("MOMO_PARTNERCODE", "ssv_momo")
// 				os.Setenv("MOMO_SECRETKEY", "1111")
// 			}
// 			momo := setupClient(c.HTTPClient)
// 			result, err := momo.CreatePayment(&c.PaymentRequest, &c.PaymentModel)
// 			assert.Equal(t, c.ExpectedQRCode, c.PaymentModel.QrCode)
// 			assert.Equal(t, c.ExpectedPaymentReq.Status, result.Status)
// 			assert.Equal(t, c.ExpectedError, err)
// 		})
// 	}
// }

// func TestGetPayment(t *testing.T) {
// 	cases := []struct {
// 		Context            string
// 		It                 string
// 		PaymentModel       model.Payment
// 		HTTPClient         httpclient.HTTPInterface
// 		ExpectedPaymentReq model.PaymentRequest
// 		ExpectedError      error
// 	}{
// 		{
// 			Context:            "Momo: test get paymStatusBadRequestent: create rsa fail",
// 			PaymentModel:       model.Payment{TransactionID: "00000", Amount: 10000},
// 			ExpectedPaymentReq: model.PaymentRequest{},
// 			ExpectedError:      errors.New("failed to parse DER encoded public key: asn1: syntax error: sequence truncated"),
// 		},
// 		{
// 			Context:            "Momo: Test get payment: call  server fail",
// 			HTTPClient:         mock.NewHTTPFailed(),
// 			PaymentModel:       model.Payment{TransactionID: "1111", Amount: 10000},
// 			ExpectedPaymentReq: model.PaymentRequest{Status: http.StatusBadRequest},
// 			ExpectedError:      errors.New("Call post request fail"),
// 		},
// 		{
// 			Context:            "Momo: Test get payment: success",
// 			HTTPClient:         mock.NewHTTPSuccess(),
// 			PaymentModel:       model.Payment{TransactionID: "111111", Amount: 10000},
// 			ExpectedPaymentReq: model.PaymentRequest{Status: http.StatusOK},
// 			ExpectedError:      nil,
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.Context, func(t *testing.T) {
// 			if c.PaymentModel.TransactionID != "00000" {
// 				os.Setenv("MOMO_PUBLICKEY", publicKey)
// 			} else {
// 				os.Setenv("MOMO_PUBLICKEY", "")
// 			}

// 			momo := setupClient(c.HTTPClient)
// 			result, err := momo.GetPayment(&c.PaymentModel)
// 			assert.Equal(t, c.ExpectedPaymentReq.Status, result.Status)
// 			assert.Equal(t, c.ExpectedError, err)
// 		})
// 	}
// }

// func TestRefundPayment(t *testing.T) {
// 	cases := []struct {
// 		Context            string
// 		It                 string
// 		PaymentModel       model.Payment
// 		HTTPClient         httpclient.HTTPInterface
// 		ExpectedPaymentReq model.PaymentRequest
// 		ExpectedError      error
// 		PaymentStatus      enum.PaymentStatus
// 	}{
// 		{
// 			Context:            "Momo: test get paymStatusBadRequestent: create rsa fail",
// 			PaymentModel:       model.Payment{TransactionID: "00000", Amount: 10000},
// 			ExpectedPaymentReq: model.PaymentRequest{},
// 			HTTPClient:         mock.NewHTTPFailed(),
// 			ExpectedError:      errors.New("failed to parse DER encoded public key: asn1: syntax error: sequence truncated"),
// 		},
// 		{
// 			Context:            "Momo: test refund payment: call server fail",
// 			HTTPClient:         mock.NewHTTPFailed(),
// 			PaymentModel:       model.Payment{TransactionID: "111111", Amount: 10000},
// 			ExpectedPaymentReq: model.PaymentRequest{Status: http.StatusBadRequest},
// 			ExpectedError:      errors.New("Call post request fail"),
// 		},
// 		{
// 			Context:            "Momo: test refund payment: success",
// 			HTTPClient:         mock.NewHTTPSuccess(),
// 			PaymentModel:       model.Payment{TransactionID: "111111", Amount: 10000},
// 			ExpectedPaymentReq: model.PaymentRequest{Status: http.StatusOK},
// 			ExpectedError:      nil,
// 			PaymentStatus:      enum.PaymentStatusRefund,
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.Context, func(t *testing.T) {
// 			if c.PaymentModel.TransactionID != "00000" {
// 				os.Setenv("MOMO_PUBLICKEY", publicKey)
// 			} else {
// 				os.Setenv("MOMO_PUBLICKEY", "")
// 			}

// 			momo := setupClient(c.HTTPClient)
// 			result, err := momo.RefundPayment(&c.PaymentModel)
// 			assert.Equal(t, c.ExpectedPaymentReq.Status, result.Status)
// 			assert.Equal(t, c.PaymentStatus, c.PaymentModel.Status)
// 			assert.Equal(t, c.ExpectedError, err)
// 		})
// 	}
// }
