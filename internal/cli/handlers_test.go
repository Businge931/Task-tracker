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

func TestHandleDelete(t *testing.T) {
	_ = os.Remove("tasks.json")
	defer os.Remove("tasks.json")

	// Setup: add a task to delete
	tasks := []task.Task{{ID: 1, Description: "To be deleted", Status: "todo", CreatedAt: "08-05-2025, 11:00am", UpdatedAt: "08-05-2025, 11:00am"}}
	_ = task.SaveTasks(tasks)

	tests := []struct {
		name        string
		args        []string
		expectError error
		shouldExist bool
	}{
		{
			name:        "valid delete",
			args:        []string{"task-cli", "delete", "1"},
			expectError: nil,
			shouldExist: false,
		},
		{
			name:        "missing args",
			args:        []string{"task-cli", "delete"},
			expectError: internalerrors.ErrUsageDelete,
			shouldExist: true,
		},
		{
			name:        "non-existent id",
			args:        []string{"task-cli", "delete", "999"},
			expectError: nil, // Will return error from DeleteTaskByID, but not usage error
			shouldExist: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := HandleDelete(tc.args)
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
					if tsk.ID == 1 {
						found = true
					}
				}
				if tc.shouldExist && !found {
					t.Errorf("task should exist but was deleted: %+v", tasks)
				}
				if !tc.shouldExist && found {
					t.Errorf("task was not deleted: %+v", tasks)
				}
			}
			// Clean up for next test
			_ = os.Remove("tasks.json")
			if tc.name != "non-existent id" {
				// Reset original task for next test
				_ = task.SaveTasks([]task.Task{{ID: 1, Description: "To be deleted", Status: "todo", CreatedAt: "08-05-2025, 11:00am", UpdatedAt: "08-05-2025, 11:00am"}})
			}
		})
	}
}

func TestHandleMarkStatus(t *testing.T) {
	_ = os.Remove("tasks.json")
	defer os.Remove("tasks.json")

	// Setup: add a task to mark
	tasks := []task.Task{{ID: 1, Description: "Mark me", Status: "todo", CreatedAt: "08-05-2025, 11:00am", UpdatedAt: "08-05-2025, 11:00am"}}
	_ = task.SaveTasks(tasks)

	tests := []struct {
		name        string
		args        []string
		expectError error
		status      string
		shouldExist bool
	}{
		{
			name:        "valid mark done",
			args:        []string{"task-cli", "mark", "1", "done"},
			expectError: nil,
			status:      "done",
			shouldExist: true,
		},
		{
			name:        "missing args",
			args:        []string{"task-cli", "mark", "1"},
			expectError: internalerrors.ErrUsageMark,
			status:      "todo",
			shouldExist: true,
		},
		{
			name:        "non-existent id",
			args:        []string{"task-cli", "mark", "999", "done"},
			expectError: nil, // Will return error from MarkTaskStatusByID, but not usage error
			status:      "",
			shouldExist: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := HandleMarkStatus(tc.args)
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
					if tsk.ID == 1 && tsk.Status == tc.status {
						found = true
					}
				}
				if tc.status != "" && !found {
					t.Errorf("task was not marked correctly: %+v", tasks)
				}
			}
			// Clean up for next test
			_ = os.Remove("tasks.json")
			if tc.name != "non-existent id" {
				// Reset original task for next test
				_ = task.SaveTasks([]task.Task{{ID: 1, Description: "Mark me", Status: "todo", CreatedAt: "08-05-2025, 11:00am", UpdatedAt: "08-05-2025, 11:00am"}})
			}
		})
	}
}

func TestHandleList(t *testing.T) {
	_ = os.Remove("tasks.json")
	defer os.Remove("tasks.json")

	tests := []struct {
		name        string
		setupTasks  []task.Task
		expectPrint string
	}{
		{
			name:        "no tasks",
			setupTasks:  nil,
			expectPrint: "No tasks found.",
		},
		{
			name:        "one task",
			setupTasks:  []task.Task{{ID: 1, Description: "Test List", Status: "todo", CreatedAt: "08-05-2025, 11:00am", UpdatedAt: "08-05-2025, 11:00am"}},
			expectPrint: "ID: 1\nDescription: Test List\nStatus: todo\nCreated: 08-05-2025, 11:00am\nUpdated: 08-05-2025, 11:00am\n---",
		},
		{
			name:        "multiple tasks",
			setupTasks:  []task.Task{{ID: 1, Description: "Task 1", Status: "todo", CreatedAt: "08-05-2025, 11:00am", UpdatedAt: "08-05-2025, 11:00am"}, {ID: 2, Description: "Task 2", Status: "done", CreatedAt: "08-05-2025, 12:00pm", UpdatedAt: "08-05-2025, 12:30pm"}},
			expectPrint: "ID: 1\nDescription: Task 1\nStatus: todo\nCreated: 08-05-2025, 11:00am\nUpdated: 08-05-2025, 11:00am\n---\nID: 2\nDescription: Task 2\nStatus: done\nCreated: 08-05-2025, 12:00pm\nUpdated: 08-05-2025, 12:30pm\n---",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_ = os.Remove("tasks.json")
			if tc.setupTasks != nil {
				_ = task.SaveTasks(tc.setupTasks)
			}

			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := HandleList([]string{"task-cli", "list"})

			w.Close()
			os.Stdout = oldStdout
			var buf [4096]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			if tc.expectPrint != "" && !containsAll(output, tc.expectPrint) {
				t.Errorf("expected output to contain:\n%s\ngot:\n%s", tc.expectPrint, output)
			}
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			_ = os.Remove("tasks.json")
		})
	}
}

// containsAll checks if all lines in expected are present in output (order matters)
func containsAll(output, expected string) bool {
	for _, line := range splitLines(expected) {
		if line == "" {
			continue
		}
		if !containsLine(output, line) {
			return false
		}
	}
	return true
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

func containsLine(output, line string) bool {
	return len(line) == 0 || (len(output) >= len(line) && contains(output, line))
}

func contains(s, substr string) bool {
	return len(substr) == 0 || (len(s) >= len(substr) && indexOf(s, substr) >= 0)
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
