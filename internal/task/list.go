package task

import "time"

func GetNextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func FormatDisplayDate(dateStr string) string {
	if t, err := time.Parse("02-01-2006, 03:04pm", dateStr); err == nil {
		return t.Format("02-01-2006, 03:04pm")
	}
	if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return t.Format("02-01-2006, 03:04pm")
	}
	return dateStr
}
