package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"uit_payment/router"
)

func main() {
	wait := time.Second * 15
	rootContext := context.Background()

	r := router.Init(rootContext)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(rootContext, wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Shutting down...")
	os.Exit(0)
}
