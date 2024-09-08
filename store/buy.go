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

func BuyItems(w http.ResponseWriter, r *http.Request) {
	_, span := telemetry.Tracer.Start(r.Context(), "buyItems")
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
	decoder := json.NewDecoder(r.Body)
	var items listItems
	err = decoder.Decode(&items)
	if err != nil {
		slog.Error(fmt.Sprintf("error to decode body request, err: %s", err.Error()))
		return
	}

	for _, item := range items.Items {
		var exist bool
		rows, err := db.Query("SELECT * FROM items WHERE name = ?", item.Name)
		if err != nil {
			slog.Error(fmt.Sprintf("error to query on the database, err: %s", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		// Loop through rows, using Scan to assign column data to struct fields.
		for rows.Next() {
			exist = true
			var i Items
			if err := rows.Scan(&i.ID, &i.Name, &i.Qtd, &i.Price); err != nil {
				slog.Error(fmt.Sprintf("error to scan row, err: %s", err.Error()))
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			newQtd := i.Qtd - item.Qtd
			if newQtd > 0 {
				_, err = db.Exec("UPDATE items SET qtd = ? WHERE name = ?", newQtd, item.Name)
				if err != nil {
					slog.Error(fmt.Sprintf("fail to add item %s on the database, err %s", item.Name, err.Error()))
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				slog.Error(fmt.Sprintf("You can't buy more items than are available in the store, item: %s", item.Name))
				http.Error(w, fmt.Sprintf("You can't buy more items than are available in the store, item: %s", item.Name), http.StatusBadRequest)
			}
		}
		if !exist {
			slog.Error(fmt.Sprintf("the item '%s' was not found in the database", item.Name))
			http.Error(w, fmt.Sprintf("the item '%s' was not found in the database", item.Name), http.StatusBadRequest)
		}
	}

}
