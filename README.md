# Kairos

A lightweight personal automation engine built in Go that lets you create, queue, and run background jobs like reminders, file tasks, and scheduled checks. It features a simple web UI and a concurrent worker system to process tasks reliably in the background.


## Features

- **Job queue** — create and manage background tasks through a simple web interface
- **Background worker** — picks up pending jobs automatically and processes them concurrently
- **Real-time status updates** — browser UI polls for job status without full page reloads (powered by htmx)
- **SQLite persistence** — zero-dependency storage, single file database

## Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- `gcc` (required by the SQLite driver)
  ```bash
  sudo apt install gcc build-essential   # Debian/Ubuntu
  ```
- [Air](https://github.com/air-verse/air) (live reload for Go)
  ```bash
  go install github.com/air-verse/air@latest
  ```

## Setup

```bash
# 1. Clone the repo
git clone https://github.com/yourname/kairos.git
cd kairos

# 2. Download the Tailwind CLI binary
make tailwind

# 3. Install Go dependencies
go mod download
```

## Development

```bash
make dev
```

This starts both the Go server (via Air) and the Tailwind CSS watcher together. Air automatically restarts the server whenever you change a `.go` or template file. The Tailwind watcher rebuilds CSS on every template save. The server runs on [http://localhost:8080](http://localhost:8080).

> **Note:** always run commands from the project root. Template and migration paths are relative to where the binary is executed from.

### Other commands

| Command | Description |
|---|---|
| `make dev` | Start server + Tailwind watcher |
| `make build` | Compile production binary with minified CSS |
| `make css` | Build minified CSS only |
| `make tailwind` | Download the Tailwind CLI binary |
| `make clean` | Remove build artifacts |

## Project Structure

```
cmd/server/         entry point
internal/
  jobs/             background worker
  storage/          database connection, migrations, job model
  web/
    static/         served at /static/ (JS, CSS)
    templates/      Go HTML templates
migrations/         SQL migration files (run in order on startup)
Makefile            developer commands
```
