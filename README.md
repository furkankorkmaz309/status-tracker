# Status Tracker

**Status Tracker** is a Go-based tracking system that periodically checks the status of services defined via JSON, saves the results in a SQLite database and sends notifications to recipients when necessary.

---

## Features

- Monitoring services by checking ping/HTTP at certain intervals
- Get data from `services.json` and `recipients.json` files
- Saving results in a SQLite database
- Support for sending notifications to recipients (notifier module)
- Configuration with .env file

---

## Installation

#### Requirements

- Go 1.21+
- `.env` file
- `services.json` and `recipients.json` files

### Steps

```bash
git clone https://github.com/furkankorkmaz309/status-tracker.git
cd status-tracker/cmd/status-tracker
go run .
````

The `.env` file needs to be added to the project root directory. Example:

``env
JSON_PATH=../../internal/assets/
DATABASE_PATH=../../internal/assets/tracker.db
```

