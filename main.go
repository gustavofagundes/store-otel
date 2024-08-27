package main

import (
	"log/slog"
	"net/http"

	"github.com/gustavofagunde/store-otel/store"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/itens", store.Itens)
	slog.Info("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		slog.Error("server stoped with error: %s", err.Error())
	}
}
