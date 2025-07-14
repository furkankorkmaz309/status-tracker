package app

import (
	"database/sql"
	"log"

	"github.com/furkankorkmaz309/status-tracker/internal/models"
)

type App struct {
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	Services   []models.Service
	Recipients []models.Recipient
	DB         *sql.DB
}
