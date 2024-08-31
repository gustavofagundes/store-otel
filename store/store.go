package store

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gustavofagunde/store-otel/db"
)

type Items struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Qtd   int64   `json:"qtd"`
	Price float32 `json:"price"`
}

func ListItems(w http.ResponseWriter, r *http.Request) {
	db, err := db.NewClient()

	if err != nil {
		slog.Error(fmt.Sprintf("fail to create client to database %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var items []Items

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		slog.Error(fmt.Sprintf("Error to encoded struct %s", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func BuyItems(w http.ResponseWriter, r *http.Request) {

}

func SaveItems(w http.ResponseWriter, r *http.Request) {

}
