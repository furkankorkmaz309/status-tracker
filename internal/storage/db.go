package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	// take path from .env
	path := os.Getenv("DB_PATH")
	if path == "" {
		return nil, fmt.Errorf("DB_PATH not set")
	}

	// create folder if not exists
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating assets folder : %v", err)
	}

	// create database
	db, err := sql.Open("sqlite3", path+"monitor.db")
	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating db file : %v", err)
	}

	// connect database
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("an error occurred while connecting to database : %v", err)
	}

	// prepare query
	query := `
	CREATE TABLE IF NOT EXISTS checks(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	service_id INTEGER,
	service_name TEXT,
	checked_at TIMESTAMP,
	status_code INTEGER
	)`

	// execute query
	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating checks table : %v", err)
	}

	// return
	return db, nil
}
