package main

import (
	"fmt"
)

func DeleteTaskByID(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("error loading tasks: %v", err)
	}
	found := false
	newTasks := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, t)
	}
	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}
	err = saveTasks(newTasks)
	if err != nil {
		return fmt.Errorf("error saving tasks: %v", err)
	}
	return nil
}
