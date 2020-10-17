package dom_api

import (
	"context"
	repository "uit_payment/entities/partner"
	"uit_payment/lib/logging"
	"uit_payment/model"

	grpc "google.golang.org/grpc"
)

type DOMService struct {
	PartnerRepo repository.PartnerRepositoryInterface
}

func NewDOMService() DOMServiceInterface {
	return &DOMService{
		PartnerRepo: repository.NewPartnerRepository(),
	}
}

func (d *DOMService) PaymentCallback(c context.Context, payment model.Payment) error {
	req := parseRequest(payment)
	partner, err := d.PartnerRepo.FindByID(payment.PartnerID)
	if err != nil {
		return err
	}

	domgRPCUrl := partner.CallbackURL
	connection, err := grpc.Dial(domgRPCUrl, grpc.WithInsecure())
	defer connection.Close()

	if err != nil {
		logging.Errorln("Did not connect to Tenant gRPC service", err.Error())
		return err
	}

	client := NewDOMServiceClient(connection)
	resp, err := client.PaymentCallback(c, &req)
	if err != nil {
		logging.Errorln("PaymentCallback:", err.Error())
		return err
	}

	logging.Println(resp.Message)

	return nil
}

func parseRequest(payment model.Payment) PaymentCallbackRequest {
	return PaymentCallbackRequest{
		TransactionId: payment.TransactionID,
		Amount:        payment.Amount,
		Status:        int32(payment.Status),
	}
}
