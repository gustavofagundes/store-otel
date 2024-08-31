package main

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavofagunde/store-otel/store"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func main() {
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("store-otel"))
	router.HandleFunc("/items", store.ListItems)
	router.HandleFunc("/buy", store.BuyItems)
	router.HandleFunc("/save", store.SaveItems)
	http.Handle("/", router)
	slog.Info("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("%s", slog.String("err", err.Error()))
	}
}
