package app

import (
	"database/sql"
	"log"
)

type App struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	DB       *sql.DB
}
