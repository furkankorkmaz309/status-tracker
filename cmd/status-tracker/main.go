package main

import (
	"log"
	"os"
	"time"

	"github.com/furkankorkmaz309/status-tracker/internal/app"
	"github.com/furkankorkmaz309/status-tracker/internal/checker"
	"github.com/furkankorkmaz309/status-tracker/internal/models"
	"github.com/furkankorkmaz309/status-tracker/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	// prepare log outputs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)

	// load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		errorLog.Fatal(err)
	}

	// load services
	serviceSlice, err := storage.LoadFile[models.Service]("JSON_PATH", "services.json")
	if err != nil {
		errorLog.Fatal(err)
	}

	// load recipients
	recipientSlice, err := storage.LoadFile[models.Recipient]("JSON_PATH", "recipients.json")
	if err != nil {
		errorLog.Fatal(err)
	}

	// initialize database
	db, err := storage.InitDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// define app with app struct
	app := &app.App{
		InfoLog:    infoLog,
		ErrorLog:   errorLog,
		Services:   serviceSlice,
		Recipients: recipientSlice,
		DB:         db,
	}

	for {
		err = checker.CheckSite(*app)
		if err != nil {
			app.ErrorLog.Fatal(err)
		}

		time.Sleep(10 * time.Minute)
	}
}
