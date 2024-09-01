package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavofagunde/store-otel/store"
	"github.com/gustavofagunde/store-otel/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func main() {
	router := mux.NewRouter()
	// Set up OpenTelemetry.
	otelShutdown, err := telemetry.SetupOTelSDK(context.TODO())
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	router.Use(otelmux.Middleware("store-otel"))
	router.HandleFunc("/items", store.ListItems)
	router.HandleFunc("/buy", store.BuyItems)
	router.HandleFunc("/add", store.AddItems)
	http.Handle("/", router)
	slog.Info("Server starting on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("%s", slog.String("err", err.Error()))
	}
}
