package notifier

import (
	"fmt"
	"os"
	"time"

	"github.com/furkankorkmaz309/status-tracker/internal/app"
	"github.com/go-gomail/gomail"
)

func SendGoMail(app app.App, msg string) error {
	// prepare mail
	header := fmt.Sprintf("Errors at %v", time.Now().Format("2006-01-02 15:04"))

	// take mail from .env
	mail := os.Getenv("MAIL")
	if mail == "" {
		return fmt.Errorf("MAIL not set")
	}

	// take password from .env
	password := os.Getenv("APP_PASSWORD")
	if password == "" {
		return fmt.Errorf("APP_PASSWORD not set")
	}

	reciepents := app.Recipients
	var emails []string
	for _, v := range reciepents {
		emails = append(emails, v.Recipient)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mail)
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", header)
	m.SetBody("text/html", msg)

	d := gomail.NewDialer("smtp.gmail.com", 587, mail, password)

	// Send mail
	err := d.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("an error occurred while sending email : %v", err)
	}

	return nil
}
