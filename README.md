
# Status Tracker

**Status Tracker** is a Go-based tracking system that periodically checks the status of services defined via JSON, saves the results in a SQLite database and sends notifications to recipients when necessary.

---

### Branch Info

- **Stable:** `main`  
- **Development:** `update-deneme`

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
```
### .env Example

Place a `.env` file in the project root with the following content:

```env
DB_PATH=../../internal/data/
JSON_PATH=../../internal/assets/
LOG_PATH=../../internal/data/
```

> Make sure the paths match your project structure.
