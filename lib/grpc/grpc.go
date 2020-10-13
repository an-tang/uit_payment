package grpc

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"

	handler "uit_payment/entities/payment/delivery/grpc"
	"uit_payment/lib/logging"
	payment_api "uit_payment/services/payment"

	"google.golang.org/grpc"
)

func RunGRPCServer(ctx context.Context, port int) error {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	paymentgRPCHandler := handler.NewPaymentgRPCHandler()
	payment_api.RegisterPaymentServiceServer(server, paymentgRPCHandler)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logging.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logging.Println("starting gRPC server...")
	return server.Serve(listen)
}
