package cli

import (
	"fmt"
	"time"

	internalerrors "github.com/Businge931/tasktracker/internal"
	"github.com/Businge931/tasktracker/internal/task"
)

func HandleAdd(args []string) error {
	if len(args) < 3 {
		return internalerrors.ErrUsageAdd
	}
	description := args[2]
	tasks, err := task.LoadTasks()
	if err != nil {
		return fmt.Errorf("error loading tasks: %w", err)
	}
	id := task.GetNextID(tasks)
	now := time.Now().Format("02-01-2006, 03:04pm")
	newTask := task.Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, newTask)
	err = task.SaveTasks(tasks)
	if err != nil {
		return fmt.Errorf("error saving tasks: %w", err)
	}
	fmt.Printf("Task added successfully (ID: %d)\n", id)
	return nil
}

func HandleUpdate(args []string) error {
	if len(args) < 4 {
		return internalerrors.ErrUsageUpdate
	}
	var id int
	_, err := fmt.Sscanf(args[2], "%d", &id)
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[2])
	}
	newDesc := args[3]
	err = task.UpdateTaskByID(id, newDesc)
	if err != nil {
		return err
	}
	fmt.Printf("Task %d updated successfully.\n", id)
	return nil
}

func HandleDelete(args []string) error {
	if len(args) < 3 {
		return internalerrors.ErrUsageDelete
	}
	var id int
	_, err := fmt.Sscanf(args[2], "%d", &id)
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[2])
	}
	err = task.DeleteTaskByID(id)
	if err != nil {
		return err
	}
	fmt.Printf("Task %d deleted successfully.\n", id)
	return nil
}

func HandleMarkStatus(args []string) error {
	if len(args) < 4 {
		return internalerrors.ErrUsageMark
	}
	var id int
	_, err := fmt.Sscanf(args[2], "%d", &id)
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[2])
	}
	status := args[3]
	err = task.MarkTaskStatusByID(id, status)
	if err != nil {
		return err
	}
	fmt.Printf("Task %d marked as %s.\n", id, status)
	return nil
}

func HandleList(args []string) error {
	tasks, err := task.LoadTasks()
	if err != nil {
		return fmt.Errorf("error loading tasks: %w", err)
	}
	if len(args) == 2 {
		// List all tasks
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}
		for _, t := range tasks {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated: %s\nUpdated: %s\n---\n", t.ID, t.Description, t.Status, task.FormatDisplayDate(t.CreatedAt), task.FormatDisplayDate(t.UpdatedAt))
		}
	} else if len(args) == 3 {
		status := args[2]
		found := false
		for _, t := range tasks {
			if t.Status == status {
				fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated: %s\nUpdated: %s\n---\n", t.ID, t.Description, t.Status, task.FormatDisplayDate(t.CreatedAt), task.FormatDisplayDate(t.UpdatedAt))
				found = true
			}
		}
		if !found {
			fmt.Printf("No tasks found with status '%s'.\n", status)
		}
	} else {
		return internalerrors.ErrUsageList
	}
	return nil
}
