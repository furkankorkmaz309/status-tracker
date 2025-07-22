package commands

import (
	"fmt"
	"net/http"
	"time"

	"github.com/furkankorkmaz309/status-tracker/internal/app"
	"github.com/furkankorkmaz309/status-tracker/internal/models"
	"github.com/furkankorkmaz309/status-tracker/internal/storage"
)

func CheckSite(app app.App) error {
	// take service slice from database
	query := `SELECT * FROM services`
	rows, err := app.DB.Query(query)
	if err != nil {
		return fmt.Errorf("an error occurred while loading services : %v", err)
	}

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err = rows.Scan(&service.ID, &service.Name, &service.Service)
		if err != nil {
			return fmt.Errorf("an error occurred while scanning service row : %v", err)
		}

		services = append(services, service)
	}

	emailStr := ""

	// check every service with for each
	// if error exists return error
	// if status code is 4xx or 5xx push to database and send email
	var count int
	var checks []models.Checks
	for _, v := range services {
		// take start time
		// start := time.Now()

		// send request
		var errMsg error
		resp, err := http.Get(v.Service)
		if err != nil {
			errMsg = fmt.Errorf("an error occurred while taking response from http: %v", err)
			app.ErrorLog.Println(errMsg)
			continue
		}
		defer resp.Body.Close()

		// calculate elapsed time
		// elapsed := time.Since(start)

		// if status code is 4xx or 5xx add to database and send email
		statusCode := resp.StatusCode
		if statusCode >= 400 {
			count++

			query := `INSERT INTO checks(service_id, service_name, checked_at, status_code) VALUES(?, ?, ?, ?)`
			_, err := app.DB.Exec(query, v.ID, v.Name, time.Now(), statusCode)

			if err != nil {
				check := models.Checks{
					ServiceID:   v.ID,
					ServiceName: v.Name,
					CheckedAt:   time.Now(),
					StatusCode:  statusCode,
				}

				checks = append(checks, check)
			}

			// fmt.Printf("Service : %v\tCode: %v\t in %v ms\n", v.Service, statusCode, float64(elapsed.Microseconds())/1000)
			emailStr += fmt.Sprintf("We found an error on %v at %v%v", v.Service, time.Now().Format(time.RFC1123), "<br>")
		}
	}

	// if can not add to database log the errors
	if len(checks) != 0 {
		filenameLog := fmt.Sprintf("failed-checks-db-error-%v.log", time.Now().Format("2006-01-02-15:04"))
		err := storage.SaveFile("LOG_PATH", filenameLog, checks)
		if err != nil {
			return err
		}
	}

	// send mail
	if count > 0 {
		err := SendGoMail(app, emailStr)
		if err != nil {
			return err
		}
	}

	// return
	return nil
}
