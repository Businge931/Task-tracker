package cli

import (
	"os"
	"testing"

	internalerrors "github.com/Businge931/tasktracker/internal"
	"github.com/Businge931/tasktracker/internal/task"
)

func TestHandleAdd(t *testing.T) {
	_ = os.Remove("tasks.json")
	defer os.Remove("tasks.json")

	tests := []struct {
		name        string
		args        []string
		expectError error
		desc        string
	}{
		{
			name:        "valid add",
			args:        []string{"task-cli", "add", "Test task from unit test"},
			expectError: nil,
			desc:        "Test task from unit test",
		},
		{
			name:        "missing description",
			args:        []string{"task-cli", "add"},
			expectError: internalerrors.ErrUsageAdd,
			desc:        "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := HandleAdd(tc.args)
			if tc.expectError != nil {
				if err == nil || err.Error() != tc.expectError.Error() {
					t.Errorf("expected error '%v', got '%v'", tc.expectError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				tasks, err := task.LoadTasks()
				if err != nil {
					t.Fatalf("failed to load tasks: %v", err)
				}
				if len(tasks) != 1 || tasks[0].Description != tc.desc {
					t.Errorf("task was not added correctly: %+v", tasks)
				}
				// Clean up for next test
				_ = os.Remove("tasks.json")
			}
		})
	}
}

func TestHandleUpdate(t *testing.T) {
	_ = os.Remove("tasks.json")
	defer os.Remove("tasks.json")

	// Setup: add a task to update
	tasks := []task.Task{{ID: 1, Description: "Original", Status: "todo", CreatedAt: "08-05-2025, 11:00am", UpdatedAt: "08-05-2025, 11:00am"}}
	_ = task.SaveTasks(tasks)

	tests := []struct {
		name        string
		args        []string
		expectError error
		newDesc     string
	}{
		{
			name:        "valid update",
			args:        []string{"task-cli", "update", "1", "Updated description"},
			expectError: nil,
			newDesc:     "Updated description",
		},
		{
			name:        "missing args",
			args:        []string{"task-cli", "update", "1"},
			expectError: internalerrors.ErrUsageUpdate,
			newDesc:     "",
		},
		{
			name:        "non-existent id",
			args:        []string{"task-cli", "update", "999", "Should not work"},
			expectError: nil, // Will return error from UpdateTaskByID, but not usage error
			newDesc:     "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := HandleUpdate(tc.args)
			if tc.expectError != nil {
				if err == nil || err.Error() != tc.expectError.Error() {
					t.Errorf("expected error '%v', got '%v'", tc.expectError, err)
				}
			} else if tc.name == "non-existent id" {
				if err == nil {
					t.Error("expected error for non-existent id, got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				tasks, err := task.LoadTasks()
				if err != nil {
					t.Fatalf("failed to load tasks: %v", err)
				}
				found := false
				for _, tsk := range tasks {
					if tsk.ID == 1 && tsk.Description == tc.newDesc {
						found = true
					}
				}
				if tc.newDesc != "" && !found {
					t.Errorf("task was not updated correctly: %+v", tasks)
				}
			}
			// Clean up for next test
			_ = os.Remove("tasks.json")
			if tc.name != "non-existent id" {
				// Reset original task for next test
				_ = task.SaveTasks([]task.Task{{ID: 1, Description: "Original", Status: "todo", CreatedAt: "08-05-2025, 11:00am", UpdatedAt: "08-05-2025, 11:00am"}})
			}
		})
	}
}
