package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func NewClient() (*sql.DB, error) {

	if db != nil {
		pingErr := db.Ping()
		if pingErr != nil {
			slog.Error(fmt.Sprintf("fail connection with the database, err: %s", pingErr.Error()))
			return nil, pingErr
		}
		return db, nil
	}
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "store",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		slog.Error(fmt.Sprintf("fail to get the handle of database connection, err: %s", err.Error()))
		return nil, err
	}

	slog.Info("Database Connected!")

	return db, nil
}
