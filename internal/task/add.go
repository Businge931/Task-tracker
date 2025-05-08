package task

import (
	"encoding/json"
	"os"
)

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(tasksFile); os.IsNotExist(err) {
		return []Task{}, nil // No file yet, return empty slice
	}
	data, err := os.ReadFile(tasksFile)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
