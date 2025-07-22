package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/furkankorkmaz309/status-tracker/internal/app"
	"github.com/furkankorkmaz309/status-tracker/internal/commands"
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

	// initialize database
	db, err := storage.InitDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// define app with app struct
	app := &app.App{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		DB:       db,
	}

	for {
		fmt.Println()
		app.InfoLog.Println("Service control process starts")

		start := time.Now()
		duration := 10 * time.Second

		app.InfoLog.Println("If you don't want to add anything, just wait for 20 seconds")
		inputService, ok := getInputWithTimeout(`If you want to add Service press "1": `, duration)

		if !ok {
			fmt.Println("\nInput timeout. Continuing...")
		} else {
			if inputService == "1" {
				var ServiceName string
				fmt.Print("Service Name : ")
				fmt.Scan(&ServiceName)

				var Service string
				fmt.Print("Service : ")
				fmt.Scan(&Service)

				err := commands.AddService(*app, ServiceName, Service)
				if err != nil {
					app.ErrorLog.Fatal(err)
				}
				app.InfoLog.Println("Service inserted successfully")
			}
		}

		inputRecipient, ok := getInputWithTimeout(`If you want to add Recipient press "2" : `, duration)

		if !ok {
			fmt.Println("\nInput timeout. Continuing...")
		} else {
			if inputRecipient == "2" {
				var Recipient string
				fmt.Print("Recipient : ")
				fmt.Scan(&Recipient)

				err = commands.AddRecipient(*app, Recipient)
				if err != nil {
					app.ErrorLog.Fatal(err)
				}
				app.InfoLog.Println("Recipient inserted successfully")
			} else {
				fmt.Println("wtf")
			}
		}

		err = commands.CheckSite(*app)
		if err != nil {
			app.ErrorLog.Fatal(err)
		}

		app.InfoLog.Println("Service control process finished. Waiting for 10 minutes")
		elapsed := time.Since(start)
		time.Sleep(10*time.Minute - elapsed)
	}
}

func getInputWithTimeout(prompt string, timeout time.Duration) (string, bool) {
	fmt.Println()
	fmt.Print(prompt)

	inputCh := make(chan string)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		inputCh <- strings.TrimSpace(text)
	}()

	select {
	case input := <-inputCh:
		return input, true
	case <-time.After(timeout):
		return "", false
	}
}
