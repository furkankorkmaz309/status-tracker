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

	// prepare checks query
	queryChecks := `
	CREATE TABLE IF NOT EXISTS checks(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	service_id INTEGER,
	service_name TEXT,
	checked_at TIMESTAMP,
	status_code INTEGER,
	FOREIGN KEY (service_id) REFERENCES services(id)
	)`

	// execute query
	_, err = db.Exec(queryChecks)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating checks table : %v", err)
	}

	// prepare service query
	queryServices := `
	CREATE TABLE IF NOT EXISTS services(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT UNIQUE,
	service TEXT UNIQUE
	)`

	// execute query
	_, err = db.Exec(queryServices)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating services table : %v", err)
	}

	// prepare recipients query
	queryRecipients := `
	CREATE TABLE IF NOT EXISTS recipients(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	recipient TEXT UNIQUE
	)`

	// execute query
	_, err = db.Exec(queryRecipients)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while creating recipient table : %v", err)
	}

	// return
	return db, nil
}
