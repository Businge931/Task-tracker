package cli

import (
	"os"
	"testing"
	"github.com/Businge931/tasktracker/internal/task"
	internalerrors "github.com/Businge931/tasktracker/internal"
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
