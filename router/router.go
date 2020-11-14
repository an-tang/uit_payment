package router

import (
	"context"

	handler "uit_payment/entities/payment/delivery/rest"
	"uit_payment/lib/logging"

	"github.com/gorilla/mux"
)

func Init(ctx context.Context) *mux.Router {
	log := logging.InitLoggerGrayLog()
	r := mux.NewRouter()

	r.Use(log.Logger)

	r.HandleFunc("/v1/payments", handler.New(handler.NewCreatePayment())).Methods("POST")
	r.HandleFunc("/v1/payments/{transaction_id}", handler.New(handler.NewGetPayment())).Methods("GET")
	r.HandleFunc("/v1/payments/{transaction_id}", handler.New(handler.NewRefundPayment())).Methods("DELETE")
	r.HandleFunc("/health", handler.New(handler.NewHealth()))

	r.HandleFunc("/momo/confirm", handler.New(handler.NewMomoConfirm())).Methods("POST")
	return r
}
