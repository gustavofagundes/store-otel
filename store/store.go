package store

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gustavofagunde/store-otel/db"
)

type Items struct {
	ID    int64
	Name  string
	Qtd   int64
	Price float32
}

func ListItems(w http.ResponseWriter, r *http.Request) {
	db, err := db.NewClient()

	if err != nil {
		slog.Error(fmt.Sprintf("fail to create client to database %s", err.Error()))
		return
	}

	var items []Items

	err = json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		slog.Error(fmt.Sprintf("error to query on the database, err: %s", err.Error()))
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item Items
		if err := rows.Scan(&item.ID, item.Name, &item.Qtd, &item.Price); err != nil {
			slog.Error(fmt.Sprintf("error to scan row, err: %s", err.Error()))
		}
		items = append(items, item)
	}

}

func BuyItems(w http.ResponseWriter, r *http.Request) {

}

func SaveItems(w http.ResponseWriter, r *http.Request) {

}
