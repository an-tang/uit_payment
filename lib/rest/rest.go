package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"uit_payment/lib/logging"
	"uit_payment/router"
)

func RunServer(port int) {
	rootContext := context.Background()

	r := router.Init(rootContext)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	logging.Println("starting REST server...")
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatalln("REST server listen: ", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Println("shutdown REST server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Fatalln("REST server shutdown:", err)
	}

	select {
	case <-ctx.Done():
		logging.Println("timeout of 1 seconds.")
	}

	logging.Println("server exiting")
}
