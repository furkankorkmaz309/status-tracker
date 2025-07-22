package commands

import (
	"fmt"

	"github.com/furkankorkmaz309/status-tracker/internal/app"
)

func AddService(app app.App, name, service string) error {
	// prepare query
	query := `INSERT OR IGNORE INTO services(name, service) VALUES(?, ?)`

	// execute query
	_, err := app.DB.Exec(query, name, service)
	if err != nil {
		return fmt.Errorf("an error occurred while inserting service : %v", err)
	}

	return nil
}
