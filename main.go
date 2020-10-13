package main

import (
	"context"

	"uit_payment/lib/grpc"
	"uit_payment/lib/rest"
)

func main() {
	go grpc.RunGRPCServer(context.TODO(), 9002)
	rest.RunServer(8081)

}
