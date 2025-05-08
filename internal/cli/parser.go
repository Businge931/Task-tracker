package cli

import (
	"fmt"
	"strings"

	"github.com/Businge931/tasktracker/internal/task"
	internalerrors "github.com/Businge931/tasktracker/internal"
)

func HandleDynamicMark(args []string) error {
	if strings.HasPrefix(args[1], "mark-") {
		markCmd := strings.TrimPrefix(args[1], "mark-")
		lastDash := strings.LastIndex(markCmd, "-")
		if lastDash == -1 || lastDash == len(markCmd)-1 {
			return internalerrors.ErrUsageMark
		}
		status := markCmd[:lastDash]
		idStr := markCmd[lastDash+1:]
		var id int
		_, err := fmt.Sscanf(idStr, "%d", &id)
		if err != nil {
			return fmt.Errorf("%w: %s", internalerrors.ErrInvalidID, idStr)
		}
		err = task.MarkTaskStatusByID(id, status)
		if err != nil {
			return err
		}
		fmt.Printf("Task %d marked as %s.\n", id, status)
		return nil
	}
	return fmt.Errorf("%w: %s", internalerrors.ErrUnknownCommand, args[1])
}
