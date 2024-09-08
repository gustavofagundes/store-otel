package store

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gustavofagunde/store-otel/db"
	"github.com/gustavofagunde/store-otel/telemetry"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func ListItems(w http.ResponseWriter, r *http.Request) {
	ctx, span := telemetry.Tracer.Start(r.Context(), "listItems")
	span.SetAttributes(
		attribute.KeyValue{Key: semconv.HTTPRequestMethodKey, Value: attribute.StringValue(r.Method)},
		attribute.KeyValue{Key: semconv.ServerAddressKey, Value: attribute.StringValue(r.URL.String())},
		attribute.KeyValue{Key: semconv.URLPathKey, Value: attribute.StringValue(r.URL.Path)},
	)
	defer span.End()
	db, err := db.NewClient()

	if err != nil {
		slog.Error(fmt.Sprintf("fail to create client to database %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var items []Items
	_, spandb := telemetry.Tracer.Start(ctx, "listItems-consultDb")
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		slog.Error(fmt.Sprintf("error to query on the database, err: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item Items
		if err := rows.Scan(&item.ID, &item.Name, &item.Qtd, &item.Price); err != nil {
			slog.Error(fmt.Sprintf("error to scan row, err: %s", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		items = append(items, item)
	}
	spandb.End()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		slog.Error(fmt.Sprintf("Error to encoded struct %s", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
