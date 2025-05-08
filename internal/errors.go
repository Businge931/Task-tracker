package internalerrors

import "errors"

var (
	ErrUsageAdd         = errors.New("Usage: task-cli add \"description\"")
	ErrUsageUpdate      = errors.New("Usage: task-cli update <id> \"new description\"")
	ErrUsageDelete      = errors.New("Usage: task-cli delete <id>")
	ErrUsageMark        = errors.New("Usage: task-cli mark-<status>-<id>")
	ErrUsageList        = errors.New("Usage: task-cli list [status]")
	ErrInvalidID        = errors.New("Invalid ID")
	ErrUnknownCommand   = errors.New("Unknown command. Available commands: add, update, delete, mark-<status>-<id>")
)
