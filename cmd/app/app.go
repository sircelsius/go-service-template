package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sircelsius/go-service-template/internal/api"
	"github.com/sircelsius/go-service-template/internal/logging"
	"github.com/sircelsius/go-service-template/internal/metrics"
	"github.com/sircelsius/go-service-template/internal/tracing"
)

// the main application.

// this should remain minimal and only import STL/internal packages, configure them and call them.
func main() {
	ctx := context.Background()


	closer := tracing.NewTracer(ctx, "my-service", true)
	defer closer.Close()

	s := api.NewServer()
	metrics.StartMetrics(ctx)

	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      s.GetRouter(),
		ReadTimeout:  time.Millisecond * 500,
		WriteTimeout: time.Millisecond * 500,
		IdleTimeout:  time.Second * 60,
	}

	logging.GetLogger(ctx).Info("starting server")

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logging.GetLogger(ctx).Error(err.Error())
		}
	}()

	logging.GetLogger(ctx).Info("server started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	logging.GetLogger(ctx).Info("shutting down")

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	srv.Shutdown(ctx)
	logging.GetLogger(ctx).Info("shutdown complete")
	os.Exit(0)
}
