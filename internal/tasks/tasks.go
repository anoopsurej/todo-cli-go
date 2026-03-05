package tasks

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/anoopsurej/todo-cli-go/store"
)

func AddTask(tasks []store.Task, description string) []store.Task {
	var id int
	if len(tasks) == 0 {
		id = 1
	} else {
		id = tasks[len(tasks)-1].ID + 1
	}
	t := store.Task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
	}

	return append(tasks, t)
}

func relativeTime(t time.Time) string {
	duration := time.Since(t)
	hours, minutes := duration.Hours(), duration.Minutes()
	if hours >= 24 {
		numDays := int(hours / 24)
		var article string
		if numDays > 1 {
			article = "days"
		} else {
			article = "day"
		}
		return fmt.Sprintf("%d %s ago", numDays, article)
	}
	if minutes >= 60 {
		numHours := int(minutes / 60)
		var article string
		if numHours > 1 {
			article = "hours"
		} else {
			article = "hour"
		}
		return fmt.Sprintf("%d %s ago", numHours, article)
	}
	if int(minutes) >= 1 {
		var article string
		if minutes > 1 {
			article = "minutes"
		} else {
			article = "minute"
		}
		return fmt.Sprintf("%d %s ago", int(minutes), article)
	}
	return "a few seconds ago"
}

func ListTasks(tasks []store.Task, all bool) {
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 5, ' ', 0)
	if all {
		fmt.Fprintln(w, "ID\tDescription\tCreated\tDone")
		for _, val := range tasks {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", val.ID, val.Description, relativeTime(val.CreatedAt), val.Completed)
		}
	} else {
		fmt.Fprintln(w, "ID\tDescription\tCreated")
		for _, val := range tasks {
			if !val.Completed {
				fmt.Fprintf(w, "%v\t%v\t%v\n", val.ID, val.Description, relativeTime(val.CreatedAt))
			}
		}
	}
	w.Flush()
}

func CompleteTask(tasks []store.Task, id int) error {
	var taskIdx int
	found := false
	for idx, val := range tasks {
		if val.ID == id {
			taskIdx = idx
			found = true
		}
	}
	if !found {
		return errors.New("Error: Task not found")
	}
	tasks[taskIdx].Completed = true
	return nil
}

func DeleteTask(tasks []store.Task, id int) ([]store.Task, error) {
	var taskIdx int
	found := false
	for idx, val := range tasks {
		if val.ID == id {
			taskIdx = idx
			found = true
		}
	}
	if !found {
		return nil, errors.New("Error: Task not found")
	}
	result := make([]store.Task, 0, len(tasks)-1)
	result = append(result, tasks[:taskIdx]...)
	result = append(result, tasks[taskIdx+1:]...)
	return result, nil
}
