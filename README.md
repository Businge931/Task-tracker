# Task Tracker CLI

A simple, maintainable, and test-driven command-line application for managing tasks, written in Go.

## Features

- **Add tasks** with descriptions
- **Update** existing tasks
- **Delete** tasks by ID
- **Mark** tasks as done or todo
- **List** all tasks with details
- Robust error handling and comprehensive unit tests

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.18 or newer installed

### Project URL

```sh
  https://roadmap.sh/projects/task-tracker
```

### Build

1. Clone this repository:
   ```sh
   git clone https://github.com/Businge931/Task-tracker.git
   cd Task-tracker
   ```
2. Build the CLI:
   ```sh
   go build -o task-cli ./cmd/task-cli
   ```

### Usage

Run the CLI with the following commands:

- **Add a task:**

  ```sh
  ./task-cli add "Your task description"
  ```

- **Update a task:**

  ```sh
  ./task-cli update <task_id> "New description"
  ```

- **Delete a task:**

  ```sh
  ./task-cli delete <task_id>
  ```

- **Mark a task as done/todo:**

  ```sh
  ./task-cli mark <task_id> done
  ./task-cli mark <task_id> todo
  ```

- **List all tasks:**
  ```sh
  ./task-cli list
  ```

### Example

```sh
./task-cli add "Buy groceries"
./task-cli list
./task-cli mark 1 done
./task-cli update 1 "Buy groceries and cook dinner"
./task-cli delete 1
```

## Running Tests

All core functionalities are covered by unit tests.

To run all tests:

```sh
go test ./internal/cli -v
```

## Project Structure

- `cmd/task-cli/main.go` — Entry point for the CLI
- `internal/cli/handlers.go` — Command handlers
- `internal/cli/handlers_test.go` — Unit tests for handlers
- `internal/task/` — Task management logic
- `internal/errors.go` — Centralized error definitions

## Notes

- Tasks are stored in a local `tasks.json` file.
- The CLI is designed for easy extensibility and robust error handling.

---

**Happy tracking!**
