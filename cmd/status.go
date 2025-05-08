package main

import (
	"fmt"
	"time"
)

func MarkTaskStatusByID(id int, status string) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("error loading tasks: %v", err)
	}
	found := false
	now := time.Now().Format("02-01-2006, 03:04pm")
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = now
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}
	err = saveTasks(tasks)
	if err != nil {
		return fmt.Errorf("error saving tasks: %v", err)
	}
	return nil
}
