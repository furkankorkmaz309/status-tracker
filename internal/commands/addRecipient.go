package commands

import (
	"fmt"

	"github.com/furkankorkmaz309/status-tracker/internal/app"
)

func AddRecipient(app app.App, recipient string) error {
	// prepare query
	query := `INSERT OR IGNORE INTO recipient(recipient) VALUES(?)`

	// execute query
	_, err := app.DB.Exec(query, recipient)
	if err != nil {
		return fmt.Errorf("an error occurred while inserting recipient : %v", err)
	}
	return nil
}
