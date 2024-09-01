package store

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gustavofagunde/store-otel/db"
)

func AddItems(w http.ResponseWriter, r *http.Request) {
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
			slog.Info(fmt.Sprintf("item %s already registred in the database, adding the quantity of the item", item.Name))
			newQtd := item.Qtd + i.Qtd
			_, err = db.Exec("UPDATE items SET qtd = ? WHERE name = ?", newQtd, item.Name)
			if err != nil {
				slog.Error(fmt.Sprintf("fail to add item %s on the database, err %s", item.Name, err.Error()))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
		if !exist {
			_, err = db.Exec("INSERT INTO items (name, qtd, price) VALUES (?, ?, ?)", item.Name, item.Qtd, item.Price)
			if err != nil {
				slog.Error(fmt.Sprintf("fail to add item %s on the database, err %s", item.Name, err.Error()))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	w.WriteHeader(http.StatusCreated)
}
